package godot

import (
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

//func NewGridMapFromPointer(ptr gdnative.Pointer) GridMap {
func newGridMapFromPointer(ptr gdnative.Pointer) GridMap {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := GridMap{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type GridMap struct {
	Spatial
	owner gdnative.Object
}

func (o *GridMap) BaseClass() string {
	return "GridMap"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *GridMap) X_UpdateOctantsCallback() {
	//log.Println("Calling GridMap.X_UpdateOctantsCallback()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "_update_octants_callback")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *GridMap) Clear() {
	//log.Println("Calling GridMap.Clear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "clear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *GridMap) ClearBakedMeshes() {
	//log.Println("Calling GridMap.ClearBakedMeshes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "clear_baked_meshes")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false idx int}], Returns: RID
*/
func (o *GridMap) GetBakeMeshInstance(idx gdnative.Int) gdnative.Rid {
	//log.Println("Calling GridMap.GetBakeMeshInstance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "get_bake_mesh_instance")

	// Call the parent method.
	// RID
	retPtr := gdnative.NewEmptyRid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRidFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Array
*/
func (o *GridMap) GetBakeMeshes() gdnative.Array {
	//log.Println("Calling GridMap.GetBakeMeshes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "get_bake_meshes")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false x int} { false y int} { false z int}], Returns: int
*/
func (o *GridMap) GetCellItem(x gdnative.Int, y gdnative.Int, z gdnative.Int) gdnative.Int {
	//log.Println("Calling GridMap.GetCellItem()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(x)
	ptrArguments[1] = gdnative.NewPointerFromInt(y)
	ptrArguments[2] = gdnative.NewPointerFromInt(z)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "get_cell_item")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false x int} { false y int} { false z int}], Returns: int
*/
func (o *GridMap) GetCellItemOrientation(x gdnative.Int, y gdnative.Int, z gdnative.Int) gdnative.Int {
	//log.Println("Calling GridMap.GetCellItemOrientation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(x)
	ptrArguments[1] = gdnative.NewPointerFromInt(y)
	ptrArguments[2] = gdnative.NewPointerFromInt(z)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "get_cell_item_orientation")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *GridMap) GetCellScale() gdnative.Real {
	//log.Println("Calling GridMap.GetCellScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "get_cell_scale")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector3
*/
func (o *GridMap) GetCellSize() gdnative.Vector3 {
	//log.Println("Calling GridMap.GetCellSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "get_cell_size")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *GridMap) GetCenterX() gdnative.Bool {
	//log.Println("Calling GridMap.GetCenterX()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "get_center_x")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *GridMap) GetCenterY() gdnative.Bool {
	//log.Println("Calling GridMap.GetCenterY()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "get_center_y")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *GridMap) GetCenterZ() gdnative.Bool {
	//log.Println("Calling GridMap.GetCenterZ()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "get_center_z")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *GridMap) GetCollisionLayer() gdnative.Int {
	//log.Println("Calling GridMap.GetCollisionLayer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "get_collision_layer")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false bit int}], Returns: bool
*/
func (o *GridMap) GetCollisionLayerBit(bit gdnative.Int) gdnative.Bool {
	//log.Println("Calling GridMap.GetCollisionLayerBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "get_collision_layer_bit")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *GridMap) GetCollisionMask() gdnative.Int {
	//log.Println("Calling GridMap.GetCollisionMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "get_collision_mask")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false bit int}], Returns: bool
*/
func (o *GridMap) GetCollisionMaskBit(bit gdnative.Int) gdnative.Bool {
	//log.Println("Calling GridMap.GetCollisionMaskBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "get_collision_mask_bit")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Array
*/
func (o *GridMap) GetMeshes() gdnative.Array {
	//log.Println("Calling GridMap.GetMeshes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "get_meshes")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *GridMap) GetOctantSize() gdnative.Int {
	//log.Println("Calling GridMap.GetOctantSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "get_octant_size")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: MeshLibrary
*/
func (o *GridMap) GetTheme() MeshLibraryImplementer {
	//log.Println("Calling GridMap.GetTheme()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "get_theme")

	// Call the parent method.
	// MeshLibrary
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newMeshLibraryFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(MeshLibraryImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "MeshLibrary" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(MeshLibraryImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: Array
*/
func (o *GridMap) GetUsedCells() gdnative.Array {
	//log.Println("Calling GridMap.GetUsedCells()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "get_used_cells")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{False true gen_lightmap_uv bool} {0.1 true lightmap_uv_texel_size float}], Returns: void
*/
func (o *GridMap) MakeBakedMeshes(genLightmapUv gdnative.Bool, lightmapUvTexelSize gdnative.Real) {
	//log.Println("Calling GridMap.MakeBakedMeshes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromBool(genLightmapUv)
	ptrArguments[1] = gdnative.NewPointerFromReal(lightmapUvTexelSize)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "make_baked_meshes")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false x int} { false y int} { false z int}], Returns: Vector3
*/
func (o *GridMap) MapToWorld(x gdnative.Int, y gdnative.Int, z gdnative.Int) gdnative.Vector3 {
	//log.Println("Calling GridMap.MapToWorld()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(x)
	ptrArguments[1] = gdnative.NewPointerFromInt(y)
	ptrArguments[2] = gdnative.NewPointerFromInt(z)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "map_to_world")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false resource Resource}], Returns: void
*/
func (o *GridMap) ResourceChanged(resource ResourceImplementer) {
	//log.Println("Calling GridMap.ResourceChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(resource.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "resource_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false x int} { false y int} { false z int} { false item int} {0 true orientation int}], Returns: void
*/
func (o *GridMap) SetCellItem(x gdnative.Int, y gdnative.Int, z gdnative.Int, item gdnative.Int, orientation gdnative.Int) {
	//log.Println("Calling GridMap.SetCellItem()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromInt(x)
	ptrArguments[1] = gdnative.NewPointerFromInt(y)
	ptrArguments[2] = gdnative.NewPointerFromInt(z)
	ptrArguments[3] = gdnative.NewPointerFromInt(item)
	ptrArguments[4] = gdnative.NewPointerFromInt(orientation)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "set_cell_item")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false scale float}], Returns: void
*/
func (o *GridMap) SetCellScale(scale gdnative.Real) {
	//log.Println("Calling GridMap.SetCellScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(scale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "set_cell_scale")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false size Vector3}], Returns: void
*/
func (o *GridMap) SetCellSize(size gdnative.Vector3) {
	//log.Println("Calling GridMap.SetCellSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "set_cell_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *GridMap) SetCenterX(enable gdnative.Bool) {
	//log.Println("Calling GridMap.SetCenterX()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "set_center_x")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *GridMap) SetCenterY(enable gdnative.Bool) {
	//log.Println("Calling GridMap.SetCenterY()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "set_center_y")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *GridMap) SetCenterZ(enable gdnative.Bool) {
	//log.Println("Calling GridMap.SetCenterZ()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "set_center_z")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool} {True true clipabove bool} {0 true floor int} {0 true axis int}], Returns: void
*/
func (o *GridMap) SetClip(enabled gdnative.Bool, clipabove gdnative.Bool, floor gdnative.Int, axis gdnative.Int) {
	//log.Println("Calling GridMap.SetClip()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)
	ptrArguments[1] = gdnative.NewPointerFromBool(clipabove)
	ptrArguments[2] = gdnative.NewPointerFromInt(floor)
	ptrArguments[3] = gdnative.NewPointerFromInt(axis)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "set_clip")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false layer int}], Returns: void
*/
func (o *GridMap) SetCollisionLayer(layer gdnative.Int) {
	//log.Println("Calling GridMap.SetCollisionLayer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(layer)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "set_collision_layer")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false bit int} { false value bool}], Returns: void
*/
func (o *GridMap) SetCollisionLayerBit(bit gdnative.Int, value gdnative.Bool) {
	//log.Println("Calling GridMap.SetCollisionLayerBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)
	ptrArguments[1] = gdnative.NewPointerFromBool(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "set_collision_layer_bit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mask int}], Returns: void
*/
func (o *GridMap) SetCollisionMask(mask gdnative.Int) {
	//log.Println("Calling GridMap.SetCollisionMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "set_collision_mask")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false bit int} { false value bool}], Returns: void
*/
func (o *GridMap) SetCollisionMaskBit(bit gdnative.Int, value gdnative.Bool) {
	//log.Println("Calling GridMap.SetCollisionMaskBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)
	ptrArguments[1] = gdnative.NewPointerFromBool(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "set_collision_mask_bit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false size int}], Returns: void
*/
func (o *GridMap) SetOctantSize(size gdnative.Int) {
	//log.Println("Calling GridMap.SetOctantSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "set_octant_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false theme MeshLibrary}], Returns: void
*/
func (o *GridMap) SetTheme(theme MeshLibraryImplementer) {
	//log.Println("Calling GridMap.SetTheme()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(theme.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "set_theme")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false pos Vector3}], Returns: Vector3
*/
func (o *GridMap) WorldToMap(pos gdnative.Vector3) gdnative.Vector3 {
	//log.Println("Calling GridMap.WorldToMap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(pos)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GridMap", "world_to_map")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

// GridMapImplementer is an interface that implements the methods
// of the GridMap class.
type GridMapImplementer interface {
	SpatialImplementer
	X_UpdateOctantsCallback()
	Clear()
	ClearBakedMeshes()
	GetBakeMeshInstance(idx gdnative.Int) gdnative.Rid
	GetBakeMeshes() gdnative.Array
	GetCellItem(x gdnative.Int, y gdnative.Int, z gdnative.Int) gdnative.Int
	GetCellItemOrientation(x gdnative.Int, y gdnative.Int, z gdnative.Int) gdnative.Int
	GetCellScale() gdnative.Real
	GetCellSize() gdnative.Vector3
	GetCenterX() gdnative.Bool
	GetCenterY() gdnative.Bool
	GetCenterZ() gdnative.Bool
	GetCollisionLayer() gdnative.Int
	GetCollisionLayerBit(bit gdnative.Int) gdnative.Bool
	GetCollisionMask() gdnative.Int
	GetCollisionMaskBit(bit gdnative.Int) gdnative.Bool
	GetMeshes() gdnative.Array
	GetOctantSize() gdnative.Int
	GetTheme() MeshLibraryImplementer
	GetUsedCells() gdnative.Array
	MakeBakedMeshes(genLightmapUv gdnative.Bool, lightmapUvTexelSize gdnative.Real)
	MapToWorld(x gdnative.Int, y gdnative.Int, z gdnative.Int) gdnative.Vector3
	ResourceChanged(resource ResourceImplementer)
	SetCellItem(x gdnative.Int, y gdnative.Int, z gdnative.Int, item gdnative.Int, orientation gdnative.Int)
	SetCellScale(scale gdnative.Real)
	SetCellSize(size gdnative.Vector3)
	SetCenterX(enable gdnative.Bool)
	SetCenterY(enable gdnative.Bool)
	SetCenterZ(enable gdnative.Bool)
	SetClip(enabled gdnative.Bool, clipabove gdnative.Bool, floor gdnative.Int, axis gdnative.Int)
	SetCollisionLayer(layer gdnative.Int)
	SetCollisionLayerBit(bit gdnative.Int, value gdnative.Bool)
	SetCollisionMask(mask gdnative.Int)
	SetCollisionMaskBit(bit gdnative.Int, value gdnative.Bool)
	SetOctantSize(size gdnative.Int)
	SetTheme(theme MeshLibraryImplementer)
	WorldToMap(pos gdnative.Vector3) gdnative.Vector3
}
