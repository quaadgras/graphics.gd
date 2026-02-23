#!/usr/bin/env python3
"""Cross-reference trivial_methods output with extension_api.json.

Usage: match.py results.jsonl extension_api.json
"""

import json
import sys


def load_results(path):
    """Load JSON lines from trivial_methods output."""
    results = {}
    with open(path) as f:
        for line in f:
            line = line.strip()
            if not line:
                continue
            obj = json.loads(line)
            key = (obj["class"], obj["method"])
            results[key] = obj
    return results


def load_bound_methods(path):
    """Extract all bound methods from extension_api.json."""
    with open(path) as f:
        api = json.load(f)

    methods = []
    for cls in api.get("classes", []):
        class_name = cls["name"]
        for m in cls.get("methods", []):
            method_name = m["name"]
            methods.append((class_name, method_name))
    return methods


def main():
    if len(sys.argv) != 3:
        print(f"Usage: {sys.argv[0]} results.jsonl extension_api.json", file=sys.stderr)
        sys.exit(1)

    results = load_results(sys.argv[1])
    bound = load_bound_methods(sys.argv[2])

    trivial = []
    non_trivial = 0
    missing = 0

    for class_name, method_name in bound:
        key = (class_name, method_name)
        if key in results:
            if results[key]["trivial"]:
                trivial.append(results[key])
            else:
                non_trivial += 1
        else:
            missing += 1

    # Sort by class then method for stable output.
    trivial.sort(key=lambda r: (r["class"], r["method"]))

    for r in trivial:
        print(f'{r["class"]}::{r["method"]}  ({r["file"]}:{r["line"]})')

    print(f"\n--- summary ---", file=sys.stderr)
    print(f"bound methods:   {len(bound)}", file=sys.stderr)
    print(f"trivial:         {len(trivial)}", file=sys.stderr)
    print(f"non-trivial:     {non_trivial}", file=sys.stderr)
    print(f"not found in TU: {missing}", file=sys.stderr)


if __name__ == "__main__":
    main()
