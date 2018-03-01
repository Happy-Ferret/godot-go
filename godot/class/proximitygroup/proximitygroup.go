package proximitygroup

import (
	"log"
	"reflect"

	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
General purpose proximity-detection node.
*/
type ProximityGroup struct {
	Spatial
}

func (o *ProximityGroup) BaseClass() string {
	return "ProximityGroup"
}

/*
   Undocumented
*/
func (o *ProximityGroup) X_ProximityGroupBroadcast(name gdnative.String, params *Variant) {
	log.Println("Calling ProximityGroup.X_ProximityGroupBroadcast()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(name)
	goArguments[1] = reflect.ValueOf(params)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_proximity_group_broadcast", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *ProximityGroup) Broadcast(name gdnative.String, parameters *Variant) {
	log.Println("Calling ProximityGroup.Broadcast()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(name)
	goArguments[1] = reflect.ValueOf(parameters)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "broadcast", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ProximityGroup) GetDispatchMode() gdnative.Int {
	log.Println("Calling ProximityGroup.GetDispatchMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_dispatch_mode", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ProximityGroup) GetGridRadius() *Vector3 {
	log.Println("Calling ProximityGroup.GetGridRadius()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_grid_radius", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ProximityGroup) GetGroupName() gdnative.String {
	log.Println("Calling ProximityGroup.GetGroupName()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_group_name", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *ProximityGroup) SetDispatchMode(mode gdnative.Int) {
	log.Println("Calling ProximityGroup.SetDispatchMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_dispatch_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ProximityGroup) SetGridRadius(radius *Vector3) {
	log.Println("Calling ProximityGroup.SetGridRadius()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(radius)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_grid_radius", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *ProximityGroup) SetGroupName(name gdnative.String) {
	log.Println("Calling ProximityGroup.SetGroupName()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_group_name", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   ProximityGroupImplementer is an interface for ProximityGroup objects.
*/
type ProximityGroupImplementer interface {
	Class
}