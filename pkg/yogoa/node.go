package yogoa

// #include <stdlib.h>
import "C"
import (
	"sync"
	"unsafe"

	"github.com/jackwakefield/yogoa/pkg/yoga"
)

type NodeType int32

const (
	NodeTypeDefault = NodeType(yoga.NodeTypeDefault)
	NodeTypeText    = NodeType(yoga.NodeTypeText)
)

type NodeMeasure func(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size
type NodeBaseline func(node *Node, width float32, height float32) float32
type NodeDirtied func(node *Node)
type NodePrint func(node *Node)

type Node struct {
	ctx    interface{}
	ref    yoga.NodeRef
	style  *NodeStyle
	layout *NodeLayout

	onMeasure  NodeMeasure
	onBaseline NodeBaseline
	onDirtied  NodeDirtied
	onPrint    NodePrint
}

var nodesByContext sync.Map

func NewNode() *Node {
	ref := yoga.NodeNew()
	return newNode(ref)
}

func NewNodeWithConfig(config *Config) *Node {
	ref := yoga.NodeNewWithConfig(config.ref)
	return newNode(ref)
}

func nodeFromRef(ref yoga.NodeRef) *Node {
	if ptr := yoga.NodeGetContext(ref); ptr != nil {
		if value, ok := nodesByContext.Load(ptr); ok {
			if node, ok := value.(*Node); ok {
				return node
			}
		}
	}
	return nil
}

func newNode(ref yoga.NodeRef) *Node {
	node := nodeFromRef(ref)
	if node == nil {
		node = &Node{
			ref: ref,
			style: &NodeStyle{
				ref: ref,
			},
			layout: &NodeLayout{
				ref: ref,
			},
		}

		var ctx unsafe.Pointer = C.malloc(C.size_t(1))
		yoga.NodeSetContext(ref, ctx)
		nodesByContext.Store(ctx, node)
	}
	return node
}

func NodeCount() int {
	return int(yoga.NodeGetInstanceCount())
}

func (n *Node) Free() {
	if n.ref != nil {
		if ptr := yoga.NodeGetContext(n.ref); ptr != nil {
			if _, ok := nodesByContext.Load(ptr); ok {
				nodesByContext.Delete(ptr)
				C.free(ptr)
			}
		}

		yoga.NodeFree(n.ref)

		n.style.ref = nil
		n.layout.ref = nil
		n.ref = nil
	}
}

func (n *Node) Clone() *Node {
	if n.ref != nil {
		return newNode(yoga.NodeClone(n.ref))
	}
	return nil
}

func (n *Node) Reset() {
	if n.ref != nil {
		yoga.NodeReset(n.ref)
	}
}

func (n *Node) Context() interface{} {
	if n.ref != nil {
		return n.ctx
	}
	return nil
}

func (n *Node) SetContext(value interface{}) {
	if n.ref != nil {
		n.ctx = value
	}
}

func (n *Node) InsertChild(child *Node, index uint32) {
	if n.ref != nil && child.ref != nil {
		yoga.NodeInsertChild(n.ref, child.ref, index)
	}
}

func (n *Node) InsertSharedChild(child *Node, index uint32) {
	if n.ref != nil && child.ref != nil {
		yoga.NodeInsertSharedChild(n.ref, child.ref, index)
	}
}

func (n *Node) AddChild(child *Node) {
	if n.ref != nil && child.ref != nil {
		yoga.NodeInsertChild(n.ref, child.ref, n.ChildCount())
	}
}

func (n *Node) RemoveChild(child *Node) {
	if n.ref != nil && child.ref != nil {
		yoga.NodeRemoveChild(n.ref, child.ref)
	}
}

func (n *Node) RemoveAllChildren() {
	if n.ref != nil {
		yoga.NodeRemoveAllChildren(n.ref)
	}
}

func (n *Node) Child(index uint32) *Node {
	if n.ref != nil {
		if child := yoga.NodeGetChild(n.ref, index); child != nil {
			return &Node{
				ref: child,
			}
		}
	}
	return nil
}

func (n *Node) Owner() *Node {
	if n.ref != nil {
		if owner := yoga.NodeGetOwner(n.ref); owner != nil {
			return &Node{
				ref: owner,
			}
		}
	}
	return nil
}

func (n *Node) ChildCount() uint32 {
	if n.ref != nil {
		return yoga.NodeGetChildCount(n.ref)
	}
	return 0
}

func (n *Node) SetChildren(children []*Node, count uint32) {
	if n.ref != nil {
		refs := make([]yoga.NodeRef, len(children))
		for i, child := range children {
			refs[i] = child.ref
		}

		yoga.NodeSetChildren(n.ref, refs, count)
	}
}

func (n *Node) CalculateLayout(availableWidth, availableHeight float32, parentDirection Direction) {
	if n.ref != nil {
		yoga.NodeCalculateLayout(n.ref, availableWidth, availableHeight, yoga.Direction(parentDirection))
	}
}

func (n *Node) MarkDirty() {
	if n.ref != nil {
		yoga.NodeMarkDirty(n.ref)
	}
}

func (n *Node) MarkDirtyRecursive() {
	if n.ref != nil {
		yoga.NodeMarkDirtyAndPropogateToDescendants(n.ref)
	}
}

func (n *Node) Print(options PrintOption) {
	if n.ref != nil {
		yoga.NodePrint(n.ref, yoga.PrintOptions(options))
	}
}

func (n *Node) CopyStyle(dest *Node) {
	if n.ref != nil && dest.ref != nil {
		yoga.NodeCopyStyle(dest.ref, n.ref)
	}
}

func (n *Node) MeasureListener() NodeMeasure {
	return n.onMeasure
}

func (n *Node) SetMeasureListener(listener NodeMeasure) {
	if n.ref != nil {
		n.onMeasure = listener

		if listener != nil {
			yoga.NodeSetMeasureFunc(n.ref, onNodeMeasure)
		}
	}
}

func onNodeMeasure(ref yoga.NodeRef, width float32, widthMode yoga.MeasureMode, height float32,
	heightMode yoga.MeasureMode) yoga.Size {
	size := yoga.Size{}

	node := nodeFromRef(ref)
	if node != nil && node.onMeasure != nil {
		result := node.onMeasure(node, width, MeasureMode(widthMode), height, MeasureMode(heightMode))
		size.Width = result.Width
		size.Height = result.Height
	}

	return size
}

func (n *Node) BaselineListener() NodeBaseline {
	return n.onBaseline
}

func (n *Node) SetBaselineListener(listener NodeBaseline) {
	if n.ref != nil {
		n.onBaseline = listener

		if listener != nil {
			yoga.NodeSetBaselineFunc(n.ref, onNodeBaseline)
		}
	}
}

func onNodeBaseline(ref yoga.NodeRef, width float32, height float32) float32 {
	node := nodeFromRef(ref)
	if node != nil && node.onBaseline != nil {
		return node.onBaseline(node, width, height)
	}
	return Undefined
}

func (n *Node) DirtiedListener() NodeDirtied {
	return n.onDirtied
}

func (n *Node) SetDirtiedListener(listener NodeDirtied) {
	if n.ref != nil {
		n.onDirtied = listener

		if listener != nil {
			yoga.NodeSetDirtiedFunc(n.ref, onNodeDirtied)
		}
	}
}

func onNodeDirtied(ref yoga.NodeRef) {
	node := nodeFromRef(ref)
	if node != nil && node.onDirtied != nil {
		node.onDirtied(node)
	}
}

func (n *Node) PrintListener() NodePrint {
	return n.onPrint
}

func (n *Node) SetPrintListener(listener NodePrint) {
	if n.ref != nil {
		n.onPrint = listener

		if listener != nil {
			yoga.NodeSetPrintFunc(n.ref, onNodePrint)
		}
	}
}

func onNodePrint(ref yoga.NodeRef) {
	node := nodeFromRef(ref)
	if node != nil && node.onPrint != nil {
		node.onPrint(node)
	}
}

func (n *Node) HasNewLayout() bool {
	if n.ref != nil {
		return yoga.NodeGetHasNewLayout(n.ref)
	}
	return false
}

func (n *Node) SetHasNewLayout(hasNewLayout bool) {
	if n.ref != nil {
		yoga.NodeSetHasNewLayout(n.ref, hasNewLayout)
	}
}

func (n *Node) NodeType() NodeType {
	if n.ref != nil {
		return NodeType(yoga.NodeGetNodeType(n.ref))
	}
	return NodeTypeDefault
}

func (n *Node) SetNodeType(nodeType NodeType) {
	if n.ref != nil {
		yoga.NodeSetNodeType(n.ref, yoga.NodeType(nodeType))
	}
}

func (n *Node) IsDirty() bool {
	if n.ref != nil {
		return yoga.NodeIsDirty(n.ref)
	}
	return false
}

func (n *Node) UsedLegacyFlag() bool {
	if n.ref != nil {
		return yoga.NodeLayoutGetDidUseLegacyFlag(n.ref)
	}
	return false
}

func (n *Node) Style() *NodeStyle {
	return n.style
}

func (n *Node) Layout() *NodeLayout {
	return n.layout
}

type NodeStyle struct {
	ref yoga.NodeRef
}

func (s *NodeStyle) Direction() Direction {
	if s.ref != nil {
		return Direction(yoga.NodeStyleGetDirection(s.ref))
	}
	return DirectionInherit
}

func (s *NodeStyle) SetDirection(direction Direction) {
	if s.ref != nil {
		yoga.NodeStyleSetDirection(s.ref, yoga.Direction(direction))
	}
}

func (s *NodeStyle) FlexDirection() FlexDirection {
	if s.ref != nil {
		return FlexDirection(yoga.NodeStyleGetFlexDirection(s.ref))
	}
	return FlexDirectionColumn
}

func (s *NodeStyle) SetFlexDirection(flexDirection FlexDirection) {
	if s.ref != nil {
		yoga.NodeStyleSetFlexDirection(s.ref, yoga.FlexDirection(flexDirection))
	}
}

func (s *NodeStyle) JustifyContent() Justify {
	if s.ref != nil {
		return Justify(yoga.NodeStyleGetJustifyContent(s.ref))
	}
	return JustifyFlexStart
}

func (s *NodeStyle) SetJustifyContent(justifyContent Justify) {
	if s.ref != nil {
		yoga.NodeStyleSetJustifyContent(s.ref, yoga.Justify(justifyContent))
	}
}

func (s *NodeStyle) AlignContent() Align {
	if s.ref != nil {
		return Align(yoga.NodeStyleGetAlignContent(s.ref))
	}
	return AlignAuto
}

func (s *NodeStyle) SetAlignContent(alignContent Align) {
	if s.ref != nil {
		yoga.NodeStyleSetAlignContent(s.ref, yoga.Align(alignContent))
	}
}

func (s *NodeStyle) AlignItems() Align {
	if s.ref != nil {
		return Align(yoga.NodeStyleGetAlignItems(s.ref))
	}
	return AlignAuto
}

func (s *NodeStyle) SetAlignItems(alignItems Align) {
	if s.ref != nil {
		yoga.NodeStyleSetAlignItems(s.ref, yoga.Align(alignItems))
	}
}

func (s *NodeStyle) AlignSelf() Align {
	if s.ref != nil {
		return Align(yoga.NodeStyleGetAlignSelf(s.ref))
	}
	return AlignAuto
}

func (s *NodeStyle) SetAlignSelf(alignSelf Align) {
	if s.ref != nil {
		yoga.NodeStyleSetAlignSelf(s.ref, yoga.Align(alignSelf))
	}
}

func (s *NodeStyle) PositionType() PositionType {
	if s.ref != nil {
		return PositionType(yoga.NodeStyleGetPositionType(s.ref))
	}
	return PositionRelative
}

func (s *NodeStyle) SetPositionType(positionType PositionType) {
	if s.ref != nil {
		yoga.NodeStyleSetPositionType(s.ref, yoga.PositionType(positionType))
	}
}

func (s *NodeStyle) FlexWrap() Wrap {
	if s.ref != nil {
		return Wrap(yoga.NodeStyleGetFlexWrap(s.ref))
	}
	return WrapNone
}

func (s *NodeStyle) SetFlexWrap(flexWrap Wrap) {
	if s.ref != nil {
		yoga.NodeStyleSetFlexWrap(s.ref, yoga.Wrap(flexWrap))
	}
}

func (s *NodeStyle) Overflow() Overflow {
	if s.ref != nil {
		return Overflow(yoga.NodeStyleGetOverflow(s.ref))
	}
	return OverflowVisible
}

func (s *NodeStyle) SetOverflow(overflow Overflow) {
	if s.ref != nil {
		yoga.NodeStyleSetOverflow(s.ref, yoga.Overflow(overflow))
	}
}

func (s *NodeStyle) Display() Display {
	if s.ref != nil {
		return Display(yoga.NodeStyleGetDisplay(s.ref))
	}
	return DisplayFlex
}

func (s *NodeStyle) SetDisplay(display Display) {
	if s.ref != nil {
		yoga.NodeStyleSetDisplay(s.ref, yoga.Display(display))
	}
}

func (s *NodeStyle) Flex() float32 {
	if s.ref != nil {
		return yoga.NodeStyleGetFlex(s.ref)
	}
	return Undefined
}

func (s *NodeStyle) SetFlex(flex float32) {
	if s.ref != nil {
		yoga.NodeStyleSetFlex(s.ref, flex)
	}
}

func (s *NodeStyle) FlexGrow() float32 {
	if s.ref != nil {
		return yoga.NodeStyleGetFlexGrow(s.ref)
	}
	return Undefined
}

func (s *NodeStyle) SetFlexGrow(flexGrow float32) {
	if s.ref != nil {
		yoga.NodeStyleSetFlexGrow(s.ref, flexGrow)
	}
}

func (s *NodeStyle) FlexShrink() float32 {
	if s.ref != nil {
		return yoga.NodeStyleGetFlexShrink(s.ref)
	}
	return Undefined
}

func (s *NodeStyle) SetFlexShrink(flexShrink float32) {
	if s.ref != nil {
		yoga.NodeStyleSetFlexShrink(s.ref, flexShrink)
	}
}

func (s *NodeStyle) FlexBasis() (flexBasis float32, unit Unit) {
	if s.ref != nil {
		value := yoga.NodeStyleGetFlexBasis(s.ref)
		return value.Value, Unit(value.Unit)
	}
	return Undefined, UnitUndefined
}

func (s *NodeStyle) SetFlexBasis(flexBasis float32) {
	if s.ref != nil {
		yoga.NodeStyleSetFlexBasis(s.ref, flexBasis)
	}
}

func (s *NodeStyle) SetFlexBasisAuto() {
	if s.ref != nil {
		yoga.NodeStyleSetFlexBasisAuto(s.ref)
	}
}

func (s *NodeStyle) SetFlexBasisPercent(flexBasis float32) {
	if s.ref != nil {
		yoga.NodeStyleSetFlexBasisPercent(s.ref, flexBasis)
	}
}

func (s *NodeStyle) Position(edge Edge) (position float32, unit Unit) {
	if s.ref != nil {
		value := yoga.NodeStyleGetPosition(s.ref, yoga.Edge(edge))
		return value.Value, Unit(value.Unit)
	}
	return Undefined, UnitUndefined
}

func (s *NodeStyle) SetPosition(edge Edge, position float32) {
	if s.ref != nil {
		yoga.NodeStyleSetPosition(s.ref, yoga.Edge(edge), position)
	}
}

func (s *NodeStyle) SetPositionPercent(edge Edge, position float32) {
	if s.ref != nil {
		yoga.NodeStyleSetPositionPercent(s.ref, yoga.Edge(edge), position)
	}
}

func (s *NodeStyle) Margin(edge Edge) (margin float32, unit Unit) {
	if s.ref != nil {
		value := yoga.NodeStyleGetMargin(s.ref, yoga.Edge(edge))
		return value.Value, Unit(value.Unit)
	}
	return Undefined, UnitUndefined
}

func (s *NodeStyle) SetMargin(edge Edge, margin float32) {
	if s.ref != nil {
		yoga.NodeStyleSetMargin(s.ref, yoga.Edge(edge), margin)
	}
}

func (s *NodeStyle) SetMarginAuto(edge Edge) {
	if s.ref != nil {
		yoga.NodeStyleSetMarginAuto(s.ref, yoga.Edge(edge))
	}
}

func (s *NodeStyle) SetMarginPercent(edge Edge, margin float32) {
	if s.ref != nil {
		yoga.NodeStyleSetMarginPercent(s.ref, yoga.Edge(edge), margin)
	}
}

func (s *NodeStyle) Padding(edge Edge) (padding float32, unit Unit) {
	if s.ref != nil {
		value := yoga.NodeStyleGetPadding(s.ref, yoga.Edge(edge))
		return value.Value, Unit(value.Unit)
	}
	return Undefined, UnitUndefined
}

func (s *NodeStyle) SetPadding(edge Edge, padding float32) {
	if s.ref != nil {
		yoga.NodeStyleSetPadding(s.ref, yoga.Edge(edge), padding)
	}
}

func (s *NodeStyle) SetPaddingPercent(edge Edge, padding float32) {
	if s.ref != nil {
		yoga.NodeStyleSetPaddingPercent(s.ref, yoga.Edge(edge), padding)
	}
}

func (s *NodeStyle) Border(edge Edge) float32 {
	if s.ref != nil {
		return yoga.NodeStyleGetBorder(s.ref, yoga.Edge(edge))
	}
	return Undefined
}

func (s *NodeStyle) SetBorder(edge Edge, border float32) {
	if s.ref != nil {
		yoga.NodeStyleSetBorder(s.ref, yoga.Edge(edge), border)
	}
}

func (s *NodeStyle) Width() (width float32, unit Unit) {
	if s.ref != nil {
		value := yoga.NodeStyleGetWidth(s.ref)
		return value.Value, Unit(value.Unit)
	}
	return Undefined, UnitUndefined
}

func (s *NodeStyle) SetWidth(width float32) {
	if s.ref != nil {
		yoga.NodeStyleSetWidth(s.ref, width)
	}
}

func (s *NodeStyle) SetWidthAuto() {
	if s.ref != nil {
		yoga.NodeStyleSetWidthAuto(s.ref)
	}
}

func (s *NodeStyle) SetWidthPercent(width float32) {
	if s.ref != nil {
		yoga.NodeStyleSetWidthPercent(s.ref, width)
	}
}

func (s *NodeStyle) Height() (height float32, unit Unit) {
	if s.ref != nil {
		value := yoga.NodeStyleGetHeight(s.ref)
		return value.Value, Unit(value.Unit)
	}
	return Undefined, UnitUndefined
}

func (s *NodeStyle) SetHeight(height float32) {
	if s.ref != nil {
		yoga.NodeStyleSetHeight(s.ref, height)
	}
}

func (s *NodeStyle) SetHeightAuto() {
	if s.ref != nil {
		yoga.NodeStyleSetHeightAuto(s.ref)
	}
}

func (s *NodeStyle) SetHeightPercent(height float32) {
	if s.ref != nil {
		yoga.NodeStyleSetHeightPercent(s.ref, height)
	}
}

func (s *NodeStyle) MinWidth() (minWidth float32, unit Unit) {
	if s.ref != nil {
		value := yoga.NodeStyleGetMinWidth(s.ref)
		return value.Value, Unit(value.Unit)
	}
	return Undefined, UnitUndefined
}

func (s *NodeStyle) SetMinWidth(minWidth float32) {
	if s.ref != nil {
		yoga.NodeStyleSetMinWidth(s.ref, minWidth)
	}
}

func (s *NodeStyle) SetMinWidthPercent(minWidth float32) {
	if s.ref != nil {
		yoga.NodeStyleSetMinWidthPercent(s.ref, minWidth)
	}
}

func (s *NodeStyle) MinHeight() (minHeight float32, unit Unit) {
	if s.ref != nil {
		value := yoga.NodeStyleGetMinHeight(s.ref)
		return value.Value, Unit(value.Unit)
	}
	return Undefined, UnitUndefined
}

func (s *NodeStyle) SetMinHeight(minHeight float32) {
	if s.ref != nil {
		yoga.NodeStyleSetMinHeight(s.ref, minHeight)
	}
}

func (s *NodeStyle) SetMinHeightPercent(minHeight float32) {
	if s.ref != nil {
		yoga.NodeStyleSetMinHeightPercent(s.ref, minHeight)
	}
}

func (s *NodeStyle) MaxWidth() (maxWidth float32, unit Unit) {
	if s.ref != nil {
		value := yoga.NodeStyleGetMaxWidth(s.ref)
		return value.Value, Unit(value.Unit)
	}
	return Undefined, UnitUndefined
}

func (s *NodeStyle) SetMaxWidth(maxWidth float32) {
	if s.ref != nil {
		yoga.NodeStyleSetMaxWidth(s.ref, maxWidth)
	}
}

func (s *NodeStyle) SetMaxWidthPercent(maxWidth float32) {
	if s.ref != nil {
		yoga.NodeStyleSetMaxWidthPercent(s.ref, maxWidth)
	}
}

func (s *NodeStyle) MaxHeight() (maxHeight float32, unit Unit) {
	if s.ref != nil {
		value := yoga.NodeStyleGetMaxHeight(s.ref)
		return value.Value, Unit(value.Unit)
	}
	return Undefined, UnitUndefined
}

func (s *NodeStyle) SetMaxHeight(maxHeight float32) {
	if s.ref != nil {
		yoga.NodeStyleSetMaxHeight(s.ref, maxHeight)
	}
}

func (s *NodeStyle) SetMaxHeightPercent(maxHeight float32) {
	if s.ref != nil {
		yoga.NodeStyleSetMaxHeightPercent(s.ref, maxHeight)
	}
}

func (s *NodeStyle) AspectRatio() float32 {
	if s.ref != nil {
		return yoga.NodeStyleGetAspectRatio(s.ref)
	}
	return Undefined
}

func (s *NodeStyle) SetAspectRatio(aspectRatio float32) {
	if s.ref != nil {
		yoga.NodeStyleSetAspectRatio(s.ref, aspectRatio)
	}
}

type NodeLayout struct {
	ref yoga.NodeRef
}

func (l *NodeLayout) Left() float32 {
	if l.ref != nil {
		return yoga.NodeLayoutGetLeft(l.ref)
	}
	return Undefined
}

func (l *NodeLayout) Top() float32 {
	if l.ref != nil {
		return yoga.NodeLayoutGetTop(l.ref)
	}
	return Undefined
}

func (l *NodeLayout) Right() float32 {
	if l.ref != nil {
		return yoga.NodeLayoutGetRight(l.ref)
	}
	return Undefined
}

func (l *NodeLayout) Bottom() float32 {
	if l.ref != nil {
		return yoga.NodeLayoutGetBottom(l.ref)
	}
	return Undefined
}

func (l *NodeLayout) Width() float32 {
	if l.ref != nil {
		return yoga.NodeLayoutGetWidth(l.ref)
	}
	return Undefined
}

func (l *NodeLayout) Height() float32 {
	if l.ref != nil {
		return yoga.NodeLayoutGetHeight(l.ref)
	}
	return Undefined
}

func (l *NodeLayout) Direction() Direction {
	if l.ref != nil {
		return Direction(yoga.NodeLayoutGetDirection(l.ref))
	}
	return DirectionInherit
}

func (l *NodeLayout) HadOverflow() bool {
	if l.ref != nil {
		return yoga.NodeLayoutGetHadOverflow(l.ref)
	}
	return false
}

func (l *NodeLayout) DidLegacyStretchFlagAffectLayout() bool {
	if l.ref != nil {
		return yoga.NodeLayoutGetDidLegacyStretchFlagAffectLayout(l.ref)
	}
	return false
}

func (l *NodeLayout) Margin(edge Edge) float32 {
	if l.ref != nil {
		return yoga.NodeLayoutGetMargin(l.ref, yoga.Edge(edge))
	}
	return Undefined
}

func (l *NodeLayout) Border(edge Edge) float32 {
	if l.ref != nil {
		return yoga.NodeLayoutGetBorder(l.ref, yoga.Edge(edge))
	}
	return Undefined
}

func (l *NodeLayout) Padding(edge Edge) float32 {
	if l.ref != nil {
		return yoga.NodeLayoutGetPadding(l.ref, yoga.Edge(edge))
	}
	return Undefined
}
