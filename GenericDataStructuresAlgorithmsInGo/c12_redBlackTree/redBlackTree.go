package redblacktree

import (
	"algorithms/GenericDataStructuresAlgorithmsInGo/model"
	"image/color"
	"log"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"github.com/mitchellh/go-homedir"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// TODO: Implementation deletion in Chapter 13 (page 545) of my book, Modern Software Development Using C#.Net, Thompson, 2006

// like AVL trees, are self-balancing
// BUT generally involve fewer rotational corrections, but the resulting tree is less balanced than an AVL tree (the depth may more than 1)

// In applications that expect many insertions and deletions and fewer searche

/*

golang lib: https://pkg.go.dev/github.com/emirpasic/gods/trees/redblacktree

https://www.youtube.com/watch?v=qvZGUFHWChY
https://www.geeksforgeeks.org/red-black-tree-set-1-introduction-2/

Simulator : https://www.cs.usfca.edu/~galles/visualization/RedBlack.html
or : https://yongdanielliang.github.io/animation/web/RBTree.html

A binary search tree is a red-black tree if
1. Every node is assigned a color of red or black.
2. The root node is always black.
3. The children of a red node are black.
4. Every path from the root node to a leaf node contains the same
number of black nodes
// nil leave if black
// new element is always inserted with a red colour

*/

// Insert https://www.geeksforgeeks.org/red-black-tree-set-2-insert/

// Insert red

// if root -> black
// if not check parent color (parent black -> keep. Parent red -> check uncle)

// if Uncle red -> change uncle and parent to black, recursive up to grandFather.
// if not (uncle black) -> rotation 4 cases -> recolor

type Node[T model.OrderedStringer] struct {
	value  T
	red    bool
	parent *Node[T]
	left   *Node[T]
	right  *Node[T]
}

type NodeDirection string

const (
	L NodeDirection = "L"
	R NodeDirection = "R"
)

type NodeRotation string

const (
	LL NodeRotation = "LL"
	RR NodeRotation = "RR"
	LR NodeRotation = "LR"
	RL NodeRotation = "RL"
)

type RedBlackTree[T model.OrderedStringer] struct {
	count int
	root  *Node[T]
}

func NewRedBlackTree[T model.OrderedStringer](value T) *RedBlackTree[T] {
	return &RedBlackTree[T]{1, &Node[T]{value, false, nil, nil, nil}}
}

func (tree *RedBlackTree[T]) Insert(val T) {
	if tree.root == nil {
		tree.root = &Node[T]{val, false, nil, nil, nil}
		tree.count++
		return
	}

	parent, nodeDirection := tree.findParent(val)
	if nodeDirection == "" {
		return
	}
	new := &Node[T]{val, true, parent, nil, nil}
	if nodeDirection == string(L) {
		parent.left = new
	} else {
		parent.right = new
	}
	tree.checkReconfigure(new)
	tree.count += 1
}

func (tree *RedBlackTree[T]) IsPresent(value T, node *Node[T]) bool {
	if node == nil {
		return false
	}
	if value < node.value {
		return tree.IsPresent(value, node.left)
	}
	if value > node.value {
		return tree.IsPresent(value, node.right)
	}
	return true
}

func (tree *RedBlackTree[T]) findParent(value T) (*Node[T], string) {
	return search(value, tree.root)
}

func (tree *RedBlackTree[T]) checkReconfigure(node *Node[T]) {
	var nodeDirection, parentDirection, rotation string
	var uncle *Node[T]
	parent := node.parent
	value := node.value
	if parent == nil || parent.parent == nil ||
		!node.red || !parent.red {
		return
	}
	grandfather := parent.parent
	if value < parent.value {
		nodeDirection = string(L)
	} else {
		nodeDirection = string(R)
	}
	if grandfather.value > parent.value {
		parentDirection = string(L)
	} else {
		parentDirection = string(R)
	}
	if parentDirection == string(L) {
		uncle = grandfather.right
	} else {
		uncle = grandfather.left
	}
	rotation = nodeDirection + parentDirection
	if uncle == nil || !uncle.red {
		if rotation == string(LL) {
			tree.rightRotate(node, parent, grandfather, true)
		} else if rotation == string(RR) {
			tree.leftRotate(node, parent, grandfather, true)
		} else if rotation == string(LR) {
			tree.rightRotate(nil, node, parent, false)
			tree.leftRotate(parent, node, grandfather, true)
			node, parent = parent, node
		} else if rotation == string(RL) {
			tree.leftRotate(nil, node, parent, false)
			tree.rightRotate(parent, node, grandfather, true)
		}
	} else {
		tree.modifyColor(grandfather)
	}
}

func (tree *RedBlackTree[T]) leftRotate(node, parent, grandfather *Node[T], modifyColor bool) {
	greatgrandfather := grandfather.parent
	tree.updateParent(parent, grandfather, greatgrandfather)
	oldLeft := parent.left
	parent.left = grandfather
	grandfather.parent = parent
	grandfather.right = oldLeft
	if oldLeft != nil {
		oldLeft.parent = grandfather
	}
	if modifyColor {
		parent.red = false
		node.red = true
		grandfather.red = true
	}
}

func (tree *RedBlackTree[T]) rightRotate(node, parent, grandfather *Node[T], modifyColor bool) {
	greatgrandfather := grandfather.parent
	tree.updateParent(parent, grandfather,
		greatgrandfather)
	oldRight := parent.right
	parent.right = grandfather
	grandfather.parent = parent
	grandfather.left = oldRight
	if oldRight != nil {
		oldRight.parent = grandfather
	}
	if modifyColor {
		parent.red = false
		node.red = true
		grandfather.red = true
	}
}

func (tree *RedBlackTree[T]) modifyColor(grandfather *Node[T]) {
	grandfather.right.red = false
	grandfather.left.red = false
	if grandfather != tree.root {
		grandfather.red = true
	}
	tree.checkReconfigure(grandfather)
}

func (tree *RedBlackTree[T]) updateParent(node, parentOldChild, newParent *Node[T]) {
	node.parent = newParent
	if newParent != nil {
		if newParent.value > parentOldChild.value {
			newParent.left = node
		} else {
			newParent.right = node
		}
	} else {
		tree.root = node
	}
}

func search[T model.OrderedStringer](value T, node *Node[T]) (*Node[T], string) {
	if value == node.value {
		return nil, ""
	} else if value > node.value {
		if node.right == nil {
			return node, string(R)
		}
		return search(value, node.right)
	} else if value < node.value {
		if node.left == nil {
			return node, string(L)
		}
		return search(value, node.left)
	}
	return nil, ""
}

// Logic for drawing tree
type NodePair struct {
	Val1, Val2 string
}
type NodePos struct {
	Val  string
	Red  bool
	YPos int
	XPos int
}

var data []NodePos
var endPoints []NodePair // Used to plot lines
func PrepareDrawTree[T model.OrderedStringer](tree RedBlackTree[T]) {
	prepareToDraw(tree)
}
func FindXY(val interface{}) (int, int) {
	for i := 0; i < len(data); i++ {
		if data[i].Val == val {
			return data[i].XPos, data[i].YPos
		}
	}
	return -1, -1
}
func FindX(val interface{}) int {
	for i := 0; i < len(data); i++ {
		if data[i].Val == val {
			return i
		}
	}
	return -1
}
func SetXValues() {
	for index := 0; index < len(data); index++ {
		xValue := FindX(data[index].Val)
		data[index].XPos = xValue
	}
}
func prepareToDraw[T model.OrderedStringer](tree RedBlackTree[T]) {
	inorderLevel(tree.root, 1)
	SetXValues()
	getEndPoints(tree.root, nil)
}
func inorderLevel[T model.OrderedStringer](node *Node[T], level int) {
	if node != nil {
		inorderLevel(node.left, level+1)
		data = append(data,
			NodePos{node.value.String(), node.red,
				100 - level, -1})
		inorderLevel(node.right, level+1)
	}
}
func getEndPoints[T model.OrderedStringer](node *Node[T], parent *Node[T]) {
	if node != nil {
		if parent != nil {
			endPoints = append(endPoints,
				NodePair{node.value.String(),
					parent.value.String()})
		}
		getEndPoints(node.left, node)
		getEndPoints(node.right, node)
	}
}

var path string

func DrawGraph(a fyne.App, w fyne.Window) {
	image := canvas.NewImageFromResource(theme.FyneLogo())
	image = canvas.NewImageFromFile(path + "tree.png")
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)
	w.Close()
	w.Show()
}
func ShowTreeGraph[T model.OrderedStringer](myTree RedBlackTree[T]) {
	PrepareDrawTree(myTree)
	myApp := app.New()
	myWindow := myApp.NewWindow("Tree")
	myWindow.Resize(fyne.NewSize(1000, 600))
	path, _ := homedir.Dir()
	path += "/Desktop//"
	nodePts := make(plotter.XYs, myTree.count)
	for i := 0; i < len(data); i++ {
		nodePts[i].Y = float64(data[i].YPos)
		nodePts[i].X = float64(data[i].XPos)
	}
	nodePtsData := nodePts
	p := plot.New()
	p.Add(plotter.NewGrid())
	nodePoints, err := plotter.NewScatter(nodePtsData)
	if err != nil {
		log.Panic(err)
	}
	nodePoints.Shape = draw.CircleGlyph{}
	nodePoints.Color = color.RGBA{R: 255, G: 255, B: 250, A: 255} // White fill
	nodePoints.Radius = vg.Points(12)
	// Plot lines
	for index := 0; index < len(endPoints); index++ {
		val1 := endPoints[index].Val1
		x1, y1 := FindXY(val1)
		val2 := endPoints[index].Val2
		x2, y2 := FindXY(val2)
		pts := plotter.XYs{{X: float64(x1), Y: float64(y1)}, {X: float64(x2), Y: float64(y2)}}
		line, err := plotter.NewLine(pts)
		if err != nil {
			log.Panic(err)
		}
		scatter, err := plotter.NewScatter(pts)
		if err != nil {
			log.Panic(err)
		}
		p.Add(line, scatter)
	}
	p.Add(nodePoints)
	// Add Labels
	for index := 0; index < len(data); index++ {
		x := float64(data[index].XPos) - 0.10
		y := float64(data[index].YPos) - 0.02
		str := data[index].Val
		if data[index].Red {
			str += "(RED)"
		} else {
			str += "(BLACK)"
		}
		label, err :=
			plotter.NewLabels(plotter.XYLabels{
				XYs: []plotter.XY{
					{X: x, Y: y},
				},
				Labels: []string{str},
			})
		if err != nil {
			log.Fatalf("could not creates labels plotter: %+v", err)
		}
		p.Add(label)
	}
	path, _ = homedir.Dir()
	path += "/Desktop/GoDS/"
	err = p.Save(1000, 600, "tree.png")
	if err != nil {
		log.Panic(err)
	}
	DrawGraph(myApp, myWindow)
	// myWindow.ShowAndRun()
}

// Make int comply with Stringer interface
type Integer int

func (i Integer) String() string {
	return strconv.Itoa(int(i))
}
