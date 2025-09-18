/*
begin_complex_operation()
begin_multicaret_edit()
for i in range(get_caret_count()):
	if multicaret_edit_ignore_caret(i):
		continue
	# Logic here.
end_multicaret_edit()
end_complex_operation()
*/

package main

func TextEdit_BeginMulticaretEdit() {
	textEdit.BeginComplexOperation()
	textEdit.BeginMulticaretEdit()
	for i := 0; i < textEdit.GetCaretCount(); i++ {
		if textEdit.MulticaretEditIgnoreCaret(i) {
			continue
		}
		// Logic here.
	}
	textEdit.EndMulticaretEdit()
	textEdit.EndComplexOperation()
}
