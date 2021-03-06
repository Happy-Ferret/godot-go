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

//func NewBitMapFromPointer(ptr gdnative.Pointer) BitMap {
func newBitMapFromPointer(ptr gdnative.Pointer) BitMap {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := BitMap{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A two-dimensional array of boolean values, can be used to efficiently store a binary matrix (every matrix element takes only one bit) and query the values using natural cartesian coordinates.
*/
type BitMap struct {
	Resource
	owner gdnative.Object
}

func (o *BitMap) BaseClass() string {
	return "BitMap"
}

/*
        Undocumented
	Args: [], Returns: Dictionary
*/
func (o *BitMap) X_GetData() gdnative.Dictionary {
	//log.Println("Calling BitMap.X_GetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitMap", "_get_data")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false arg0 Dictionary}], Returns: void
*/
func (o *BitMap) X_SetData(arg0 gdnative.Dictionary) {
	//log.Println("Calling BitMap.X_SetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromDictionary(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitMap", "_set_data")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Creates a bitmap with the specified size, filled with false.
	Args: [{ false size Vector2}], Returns: void
*/
func (o *BitMap) Create(size gdnative.Vector2) {
	//log.Println("Calling BitMap.Create()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitMap", "create")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Creates a bitmap that matches the given image dimensions, every element of the bitmap is set to false if the alpha value of the image at that position is equal to [code]threshold[/code] or less, and true in other case.
	Args: [{ false image Image}], Returns: void
*/
func (o *BitMap) CreateFromImageAlpha(image ImageImplementer) {
	//log.Println("Calling BitMap.CreateFromImageAlpha()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(image.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitMap", "create_from_image_alpha")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns bitmap's value at the specified position.
	Args: [{ false position Vector2}], Returns: bool
*/
func (o *BitMap) GetBit(position gdnative.Vector2) gdnative.Bool {
	//log.Println("Calling BitMap.GetBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitMap", "get_bit")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns bitmap's dimensions.
	Args: [], Returns: Vector2
*/
func (o *BitMap) GetSize() gdnative.Vector2 {
	//log.Println("Calling BitMap.GetSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitMap", "get_size")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Returns the amount of bitmap elements that are set to true.
	Args: [], Returns: int
*/
func (o *BitMap) GetTrueBitCount() gdnative.Int {
	//log.Println("Calling BitMap.GetTrueBitCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitMap", "get_true_bit_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Sets the bitmap's element at the specified position, to the specified value.
	Args: [{ false position Vector2} { false bit bool}], Returns: void
*/
func (o *BitMap) SetBit(position gdnative.Vector2, bit gdnative.Bool) {
	//log.Println("Calling BitMap.SetBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromVector2(position)
	ptrArguments[1] = gdnative.NewPointerFromBool(bit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitMap", "set_bit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets a rectangular portion of the bitmap to the specified value.
	Args: [{ false p_rect Rect2} { false bit bool}], Returns: void
*/
func (o *BitMap) SetBitRect(pRect gdnative.Rect2, bit gdnative.Bool) {
	//log.Println("Calling BitMap.SetBitRect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromRect2(pRect)
	ptrArguments[1] = gdnative.NewPointerFromBool(bit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BitMap", "set_bit_rect")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// BitMapImplementer is an interface that implements the methods
// of the BitMap class.
type BitMapImplementer interface {
	ResourceImplementer
	X_GetData() gdnative.Dictionary
	X_SetData(arg0 gdnative.Dictionary)
	Create(size gdnative.Vector2)
	CreateFromImageAlpha(image ImageImplementer)
	GetBit(position gdnative.Vector2) gdnative.Bool
	GetSize() gdnative.Vector2
	GetTrueBitCount() gdnative.Int
	SetBit(position gdnative.Vector2, bit gdnative.Bool)
	SetBitRect(pRect gdnative.Rect2, bit gdnative.Bool)
}
