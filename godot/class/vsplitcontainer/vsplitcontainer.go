package vsplitcontainer

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
Vertical split container. See [SplitContainer]. This goes from left to right.
*/
type VSplitContainer struct {
	SplitContainer
}

func (o *VSplitContainer) BaseClass() string {
	return "VSplitContainer"
}

/*
   VSplitContainerImplementer is an interface for VSplitContainer objects.
*/
type VSplitContainerImplementer interface {
	Class
}