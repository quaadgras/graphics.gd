// trivial_methods: clang libtooling analyzer that classifies C++ methods
// as trivial (no function calls, no new/delete) or non-trivial.
//
// Usage: trivial_methods -p <compile_commands_dir> [source files...]
// Output: JSON lines to stdout.

#include "clang/AST/ASTConsumer.h"
#include "clang/AST/RecursiveASTVisitor.h"
#include "clang/Frontend/CompilerInstance.h"
#include "clang/Frontend/FrontendAction.h"
#include "clang/Tooling/CommonOptionsParser.h"
#include "clang/Tooling/Tooling.h"
#include "llvm/Support/CommandLine.h"
#include "llvm/Support/raw_ostream.h"

#include <mutex>
#include <set>
#include <string>

using namespace clang;
using namespace clang::tooling;

static llvm::cl::OptionCategory ToolCategory("trivial-methods options");

// Detect any expression that implies a function call or dynamic allocation.
class CallFinder : public RecursiveASTVisitor<CallFinder> {
public:
    bool Found = false;

    bool VisitCallExpr(CallExpr *) {
        Found = true;
        return false; // stop traversal
    }
    bool VisitCXXNewExpr(CXXNewExpr *) {
        Found = true;
        return false;
    }
    bool VisitCXXDeleteExpr(CXXDeleteExpr *) {
        Found = true;
        return false;
    }
};

static std::mutex OutputMu;
static std::set<std::string> Seen;

class MethodVisitor : public RecursiveASTVisitor<MethodVisitor> {
    ASTContext &Ctx;

public:
    explicit MethodVisitor(ASTContext &C) : Ctx(C) {}

    bool VisitCXXMethodDecl(CXXMethodDecl *MD) {
        // Skip implicit (compiler-generated) methods.
        if (MD->isImplicit())
            return true;

        // Skip template instantiations — only inspect primary templates.
        if (MD->getTemplateSpecializationKind() == TSK_ImplicitInstantiation)
            return true;

        // Skip constructors, destructors, conversion operators (never bound).
        if (isa<CXXConstructorDecl>(MD) || isa<CXXDestructorDecl>(MD) ||
            isa<CXXConversionDecl>(MD))
            return true;

        // Must have a body in this TU.
        if (!MD->hasBody())
            return true;

        // Must belong to a named class.
        const CXXRecordDecl *RD = MD->getParent();
        if (!RD || RD->getNameAsString().empty())
            return true;

        std::string QualName = MD->getQualifiedNameAsString();

        // Deduplicate: inline header methods appear in every TU.
        {
            std::lock_guard<std::mutex> Lock(OutputMu);
            if (!Seen.insert(QualName).second)
                return true;
        }

        // Look for calls in the body.
        CallFinder CF;
        CF.TraverseStmt(MD->getBody());

        bool Trivial = !CF.Found;

        // Get source location info.
        SourceManager &SM = Ctx.getSourceManager();
        SourceLocation Loc = MD->getLocation();
        std::string File = SM.getFilename(Loc).str();
        unsigned Line = SM.getSpellingLineNumber(Loc);

        std::string ClassName = RD->getNameAsString();
        std::string MethodName = MD->getNameAsString();

        // Escape strings for JSON (just backslash and double-quote).
        auto Escape = [](const std::string &S) -> std::string {
            std::string R;
            R.reserve(S.size());
            for (char C : S) {
                if (C == '"' || C == '\\')
                    R += '\\';
                R += C;
            }
            return R;
        };

        std::string JSON;
        llvm::raw_string_ostream OS(JSON);
        OS << "{\"class\":\"" << Escape(ClassName)
           << "\",\"method\":\"" << Escape(MethodName)
           << "\",\"file\":\"" << Escape(File)
           << "\",\"line\":" << Line
           << ",\"trivial\":" << (Trivial ? "true" : "false")
           << "}";

        {
            std::lock_guard<std::mutex> Lock(OutputMu);
            llvm::outs() << OS.str() << "\n";
        }

        return true;
    }
};

class MethodConsumer : public ASTConsumer {
    MethodVisitor Visitor;

public:
    explicit MethodConsumer(ASTContext &Ctx) : Visitor(Ctx) {}

    void HandleTranslationUnit(ASTContext &Ctx) override {
        Visitor.TraverseDecl(Ctx.getTranslationUnitDecl());
    }
};

class MethodAction : public ASTFrontendAction {
public:
    std::unique_ptr<ASTConsumer>
    CreateASTConsumer(CompilerInstance &CI, StringRef) override {
        return std::make_unique<MethodConsumer>(CI.getASTContext());
    }
};

int main(int argc, const char **argv) {
    auto Parser = CommonOptionsParser::create(argc, argv, ToolCategory);
    if (!Parser) {
        llvm::errs() << Parser.takeError();
        return 1;
    }

    ClangTool Tool(Parser->getCompilations(), Parser->getSourcePathList());
    return Tool.run(newFrontendActionFactory<MethodAction>().get());
}
