package vehicle

import "log"

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
This node needs to be used as a child node of [VehicleBody] and simulates the behaviour of one of its wheels. This node also acts as a collider to detect if the wheel is touching a surface.
*/
type VehicleWheel struct {
	spatial
}

func (o *VehicleWheel) BaseClass() string {
	return "VehicleWheel"
}

/*
   Undocumented
*/
func (o *VehicleWheel) GetDampingCompression() gdnative.Float {
	log.Println("Calling VehicleWheel.GetDampingCompression()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_damping_compression")

	// Call the parent method.
	// float
	retPtr := NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	log.Println("  Got return value: ", returnValue)
	return ret

}

/*
   Undocumented
*/
func (o *VehicleWheel) GetDampingRelaxation() gdnative.Float {
	log.Println("Calling VehicleWheel.GetDampingRelaxation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_damping_relaxation")

	// Call the parent method.
	// float
	retPtr := NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	log.Println("  Got return value: ", returnValue)
	return ret

}

/*
   Undocumented
*/
func (o *VehicleWheel) GetFrictionSlip() gdnative.Float {
	log.Println("Calling VehicleWheel.GetFrictionSlip()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_friction_slip")

	// Call the parent method.
	// float
	retPtr := NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	log.Println("  Got return value: ", returnValue)
	return ret

}

/*
   Undocumented
*/
func (o *VehicleWheel) GetRadius() gdnative.Float {
	log.Println("Calling VehicleWheel.GetRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_radius")

	// Call the parent method.
	// float
	retPtr := NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	log.Println("  Got return value: ", returnValue)
	return ret

}

/*
   Undocumented
*/
func (o *VehicleWheel) GetRollInfluence() gdnative.Float {
	log.Println("Calling VehicleWheel.GetRollInfluence()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_roll_influence")

	// Call the parent method.
	// float
	retPtr := NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	log.Println("  Got return value: ", returnValue)
	return ret

}

/*
   Returns a value between 0.0 and 1.0 that indicates whether this wheel is skidding. 0.0 is not skidding, 1.0 means the wheel has lost grip.
*/
func (o *VehicleWheel) GetSkidinfo() gdnative.Float {
	log.Println("Calling VehicleWheel.GetSkidinfo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_skidinfo")

	// Call the parent method.
	// float
	retPtr := NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	log.Println("  Got return value: ", returnValue)
	return ret

}

/*
   Undocumented
*/
func (o *VehicleWheel) GetSuspensionMaxForce() gdnative.Float {
	log.Println("Calling VehicleWheel.GetSuspensionMaxForce()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_suspension_max_force")

	// Call the parent method.
	// float
	retPtr := NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	log.Println("  Got return value: ", returnValue)
	return ret

}

/*
   Undocumented
*/
func (o *VehicleWheel) GetSuspensionRestLength() gdnative.Float {
	log.Println("Calling VehicleWheel.GetSuspensionRestLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_suspension_rest_length")

	// Call the parent method.
	// float
	retPtr := NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	log.Println("  Got return value: ", returnValue)
	return ret

}

/*
   Undocumented
*/
func (o *VehicleWheel) GetSuspensionStiffness() gdnative.Float {
	log.Println("Calling VehicleWheel.GetSuspensionStiffness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_suspension_stiffness")

	// Call the parent method.
	// float
	retPtr := NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	log.Println("  Got return value: ", returnValue)
	return ret

}

/*
   Undocumented
*/
func (o *VehicleWheel) GetSuspensionTravel() gdnative.Float {
	log.Println("Calling VehicleWheel.GetSuspensionTravel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_suspension_travel")

	// Call the parent method.
	// float
	retPtr := NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	log.Println("  Got return value: ", returnValue)
	return ret

}

/*
   Returns true if this wheel is in contact with a surface.
*/
func (o *VehicleWheel) IsInContact() gdnative.Bool {
	log.Println("Calling VehicleWheel.IsInContact()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "is_in_contact")

	// Call the parent method.
	// bool
	retPtr := NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	log.Println("  Got return value: ", returnValue)
	return ret

}

/*
   Undocumented
*/
func (o *VehicleWheel) IsUsedAsSteering() gdnative.Bool {
	log.Println("Calling VehicleWheel.IsUsedAsSteering()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "is_used_as_steering")

	// Call the parent method.
	// bool
	retPtr := NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	log.Println("  Got return value: ", returnValue)
	return ret

}

/*
   Undocumented
*/
func (o *VehicleWheel) IsUsedAsTraction() gdnative.Bool {
	log.Println("Calling VehicleWheel.IsUsedAsTraction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "is_used_as_traction")

	// Call the parent method.
	// bool
	retPtr := NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	log.Println("  Got return value: ", returnValue)
	return ret

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetDampingCompression(length gdnative.Float) {
	log.Println("Calling VehicleWheel.SetDampingCompression()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(length)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_damping_compression")

	// Call the parent method.
	// void
	retPtr := NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetDampingRelaxation(length gdnative.Float) {
	log.Println("Calling VehicleWheel.SetDampingRelaxation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(length)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_damping_relaxation")

	// Call the parent method.
	// void
	retPtr := NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetFrictionSlip(length gdnative.Float) {
	log.Println("Calling VehicleWheel.SetFrictionSlip()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(length)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_friction_slip")

	// Call the parent method.
	// void
	retPtr := NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetRadius(length gdnative.Float) {
	log.Println("Calling VehicleWheel.SetRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(length)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_radius")

	// Call the parent method.
	// void
	retPtr := NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetRollInfluence(rollInfluence gdnative.Float) {
	log.Println("Calling VehicleWheel.SetRollInfluence()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(rollInfluence)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_roll_influence")

	// Call the parent method.
	// void
	retPtr := NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetSuspensionMaxForce(length gdnative.Float) {
	log.Println("Calling VehicleWheel.SetSuspensionMaxForce()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(length)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_suspension_max_force")

	// Call the parent method.
	// void
	retPtr := NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetSuspensionRestLength(length gdnative.Float) {
	log.Println("Calling VehicleWheel.SetSuspensionRestLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(length)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_suspension_rest_length")

	// Call the parent method.
	// void
	retPtr := NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetSuspensionStiffness(length gdnative.Float) {
	log.Println("Calling VehicleWheel.SetSuspensionStiffness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(length)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_suspension_stiffness")

	// Call the parent method.
	// void
	retPtr := NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetSuspensionTravel(length gdnative.Float) {
	log.Println("Calling VehicleWheel.SetSuspensionTravel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(length)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_suspension_travel")

	// Call the parent method.
	// void
	retPtr := NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetUseAsSteering(enable gdnative.Bool) {
	log.Println("Calling VehicleWheel.SetUseAsSteering()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_use_as_steering")

	// Call the parent method.
	// void
	retPtr := NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
   Undocumented
*/
func (o *VehicleWheel) SetUseAsTraction(enable gdnative.Bool) {
	log.Println("Calling VehicleWheel.SetUseAsTraction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_use_as_traction")

	// Call the parent method.
	// void
	retPtr := NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
   VehicleWheelImplementer is an interface for VehicleWheel objects.
*/
type VehicleWheelImplementer interface {
	Class
}