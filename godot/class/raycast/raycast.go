package raycast

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
A RayCast represents a line from its origin to its destination position, [code]cast_to[/code]. It is used to query the 3D space in order to find the closest object along the path of the ray. RayCast can ignore some objects by adding them to the exception list via [code]add_exception[/code], by setting proper filtering with collision layers, or by filtering object types with type masks. Only enabled raycasts will be able to query the space and report collisions. RayCast calculates intersection every physics frame (see [Node]), and the result is cached so it can be used later until the next frame. If multiple queries are required between physics frames (or during the same frame) use [method force_raycast_update] after adjusting the raycast.
*/
type RayCast struct {
	Spatial
}

func (o *RayCast) BaseClass() string {
	return "RayCast"
}

/*
   Adds a collision exception so the ray does not report collisions with the specified node.
*/
func (o *RayCast) AddException(node *Object) {
	log.Println("Calling RayCast.AddException()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(node)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "add_exception", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Adds a collision exception so the ray does not report collisions with the specified [RID].
*/
func (o *RayCast) AddExceptionRid(rid *RID) {
	log.Println("Calling RayCast.AddExceptionRid()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(rid)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "add_exception_rid", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Removes all collision exceptions for this ray.
*/
func (o *RayCast) ClearExceptions() {
	log.Println("Calling RayCast.ClearExceptions()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "clear_exceptions", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Updates the collision information for the ray. Use this method to update the collision information immediately instead of waiting for the next [code]_physics_process[/code] call, for example if the ray or its parent has changed state. Note: [code]enabled == true[/code] is not required for this to work.
*/
func (o *RayCast) ForceRaycastUpdate() {
	log.Println("Calling RayCast.ForceRaycastUpdate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "force_raycast_update", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RayCast) GetCastTo() *Vector3 {
	log.Println("Calling RayCast.GetCastTo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_cast_to", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the closest object the ray is pointing to. Note that this does not consider the length of the ray, so you must also use [method is_colliding] to check if the object returned is actually colliding with the ray. Example: [codeblock] if RayCast.is_colliding(): var collider = RayCast.get_collider() [/codeblock]
*/
func (o *RayCast) GetCollider() *Object {
	log.Println("Calling RayCast.GetCollider()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collider", goArguments, "*Object")

	returnValue := goRet.Interface().(*Object)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the collision shape of the closest object the ray is pointing to. Note that this does not consider the length of the ray, so you must also use [method is_colliding] to check if the object returned is actually colliding with the ray. Example: [codeblock] if RayCast.is_colliding(): var shape = RayCast.get_collider_shape() [/codeblock]
*/
func (o *RayCast) GetColliderShape() gdnative.Int {
	log.Println("Calling RayCast.GetColliderShape()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collider_shape", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RayCast) GetCollisionMask() gdnative.Int {
	log.Println("Calling RayCast.GetCollisionMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collision_mask", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if the bit index passed is turned on. Note that bit indexes range from 0-19.
*/
func (o *RayCast) GetCollisionMaskBit(bit gdnative.Int) gdnative.Bool {
	log.Println("Calling RayCast.GetCollisionMaskBit()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(bit)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collision_mask_bit", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the normal of the intersecting object's shape at the collision point.
*/
func (o *RayCast) GetCollisionNormal() *Vector3 {
	log.Println("Calling RayCast.GetCollisionNormal()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collision_normal", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the collision point at which the ray intersects the closest object. Note: this point is in the [b]global[/b] coordinate system.
*/
func (o *RayCast) GetCollisionPoint() *Vector3 {
	log.Println("Calling RayCast.GetCollisionPoint()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collision_point", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RayCast) GetExcludeParentBody() gdnative.Bool {
	log.Println("Calling RayCast.GetExcludeParentBody()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_exclude_parent_body", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return whether the closest object the ray is pointing to is colliding with the vector (considering the vector length).
*/
func (o *RayCast) IsColliding() gdnative.Bool {
	log.Println("Calling RayCast.IsColliding()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_colliding", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RayCast) IsEnabled() gdnative.Bool {
	log.Println("Calling RayCast.IsEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_enabled", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Removes a collision exception so the ray does report collisions with the specified node.
*/
func (o *RayCast) RemoveException(node *Object) {
	log.Println("Calling RayCast.RemoveException()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(node)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "remove_exception", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Removes a collision exception so the ray does report collisions with the specified [RID].
*/
func (o *RayCast) RemoveExceptionRid(rid *RID) {
	log.Println("Calling RayCast.RemoveExceptionRid()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(rid)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "remove_exception_rid", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RayCast) SetCastTo(localPoint *Vector3) {
	log.Println("Calling RayCast.SetCastTo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(localPoint)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_cast_to", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RayCast) SetCollisionMask(mask gdnative.Int) {
	log.Println("Calling RayCast.SetCollisionMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mask)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_collision_mask", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the bit index passed to the [code]value[/code] passed. Note that bit indexes range from 0-19.
*/
func (o *RayCast) SetCollisionMaskBit(bit gdnative.Int, value gdnative.Bool) {
	log.Println("Calling RayCast.SetCollisionMaskBit()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(bit)
	goArguments[1] = reflect.ValueOf(value)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_collision_mask_bit", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RayCast) SetEnabled(enabled gdnative.Bool) {
	log.Println("Calling RayCast.SetEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enabled)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_enabled", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RayCast) SetExcludeParentBody(mask gdnative.Bool) {
	log.Println("Calling RayCast.SetExcludeParentBody()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mask)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_exclude_parent_body", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   RayCastImplementer is an interface for RayCast objects.
*/
type RayCastImplementer interface {
	Class
}