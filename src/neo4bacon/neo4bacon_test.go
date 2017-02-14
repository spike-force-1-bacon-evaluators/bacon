package neo4bacon

import "testing"

func TestMapResult(t *testing.T) {

	valuesA := []interface{}{"r1a", "r2a", "r3a", "r4a", "r5a", "r6a", "r7a"}
	valuesB := []interface{}{"r1b", "r2b", "r3b", "r4b", "r5b", "r6b", "r7b"}
	a := []interface{}{valuesA[0], valuesB[0]} // 1
	b := []interface{}{valuesA[1], valuesB[1]} // 2
	c := []interface{}{valuesA[2], valuesB[2]} // 3
	d := []interface{}{valuesA[3], valuesB[3]} // 4
	e := []interface{}{valuesA[1], valuesB[1]} // 2
	f := []interface{}{valuesA[0], valuesB[0]} // 1
	g := []interface{}{valuesA[2], valuesB[2]} // 3
	h := []interface{}{valuesA[4], valuesB[4]} // 3

	var newList [][]interface{}
	newList = append(newList, a, b, h, c, d) // 1, 2, 5, 3, 4

	var oldList [][]interface{}
	oldList = append(oldList, d, e, f, g) // 4, 2, 1, 3

	n := &neo4j{
		newlistresult: newList,
		oldlistresult: oldList,
	}

	n.mapResult()

	var expected [][]string

	e1 := []string{"r1a", "2"}
	e2 := []string{"r2a", "0"}
	e3 := []string{"r5a", "N/A"}
	e4 := []string{"r3a", "0"}
	e5 := []string{"r4a", "-4"}

	expected = append(expected, e1, e2, e3, e4, e5)

	for i, r := range n.result {
		if r.ID != expected[i][0] {
			t.Errorf("Expected (r.ID) %s == %s", r.ID, expected[i][0])
		}

		if r.Change != expected[i][1] {
			t.Errorf("Expected (r.Change) %s == %s", r.Change, expected[i][0])
		}
	}
}
