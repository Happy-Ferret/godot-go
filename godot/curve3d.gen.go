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

//func NewCurve3DFromPointer(ptr gdnative.Pointer) Curve3D {
func newCurve3DFromPointer(ptr gdnative.Pointer) Curve3D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Curve3D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
This class describes a Bezier curve in 3D space. It is mainly used to give a shape to a [Path], but can be manually sampled for other purposes. It keeps a cache of precalculated points along the curve, to speed further calculations up.
*/
type Curve3D struct {
	Resource
	owner gdnative.Object
}

func (o *Curve3D) BaseClass() string {
	return "Curve3D"
}

/*
        Undocumented
	Args: [], Returns: Dictionary
*/
func (o *Curve3D) X_GetData() gdnative.Dictionary {
	//log.Println("Calling Curve3D.X_GetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "_get_data")

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
func (o *Curve3D) X_SetData(arg0 gdnative.Dictionary) {
	//log.Println("Calling Curve3D.X_SetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromDictionary(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "_set_data")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a point to a curve, at "position", with control points "in" and "out". If "at_position" is given, the point is inserted before the point number "at_position", moving that point (and every point after) after the inserted point. If "at_position" is not given, or is an illegal value (at_position <0 or at_position >= [method get_point_count]), the point will be appended at the end of the point list.
	Args: [{ false position Vector3} {(0, 0, 0) true in Vector3} {(0, 0, 0) true out Vector3} {-1 true at_position int}], Returns: void
*/
func (o *Curve3D) AddPoint(position gdnative.Vector3, in gdnative.Vector3, out gdnative.Vector3, atPosition gdnative.Int) {
	//log.Println("Calling Curve3D.AddPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromVector3(position)
	ptrArguments[1] = gdnative.NewPointerFromVector3(in)
	ptrArguments[2] = gdnative.NewPointerFromVector3(out)
	ptrArguments[3] = gdnative.NewPointerFromInt(atPosition)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "add_point")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes all points from the curve.
	Args: [], Returns: void
*/
func (o *Curve3D) ClearPoints() {
	//log.Println("Calling Curve3D.ClearPoints()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "clear_points")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Curve3D) GetBakeInterval() gdnative.Float {
	//log.Println("Calling Curve3D.GetBakeInterval()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_bake_interval")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Returns the total length of the curve, based on the cached points. Given enough density (see [method set_bake_interval]), it should be approximate enough.
	Args: [], Returns: float
*/
func (o *Curve3D) GetBakedLength() gdnative.Float {
	//log.Println("Calling Curve3D.GetBakedLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_baked_length")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Returns the cache of points as a [PoolVector3Array].
	Args: [], Returns: PoolVector3Array
*/
func (o *Curve3D) GetBakedPoints() gdnative.PoolVector3Array {
	//log.Println("Calling Curve3D.GetBakedPoints()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_baked_points")

	// Call the parent method.
	// PoolVector3Array
	retPtr := gdnative.NewEmptyPoolVector3Array()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolVector3ArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the cache of tilts as a [RealArray].
	Args: [], Returns: PoolRealArray
*/
func (o *Curve3D) GetBakedTilts() gdnative.PoolRealArray {
	//log.Println("Calling Curve3D.GetBakedTilts()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_baked_tilts")

	// Call the parent method.
	// PoolRealArray
	retPtr := gdnative.NewEmptyPoolRealArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolRealArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the number of points describing the curve.
	Args: [], Returns: int
*/
func (o *Curve3D) GetPointCount() gdnative.Int {
	//log.Println("Calling Curve3D.GetPointCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_point_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the position of the control point leading to the vertex "idx". If the index is out of bounds, the function sends an error to the console, and returns (0, 0, 0).
	Args: [{ false idx int}], Returns: Vector3
*/
func (o *Curve3D) GetPointIn(idx gdnative.Int) gdnative.Vector3 {
	//log.Println("Calling Curve3D.GetPointIn()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_point_in")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns the position of the control point leading out of the vertex "idx". If the index is out of bounds, the function sends an error to the console, and returns (0, 0, 0).
	Args: [{ false idx int}], Returns: Vector3
*/
func (o *Curve3D) GetPointOut(idx gdnative.Int) gdnative.Vector3 {
	//log.Println("Calling Curve3D.GetPointOut()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_point_out")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns the position of the vertex "idx". If the index is out of bounds, the function sends an error to the console, and returns (0, 0, 0).
	Args: [{ false idx int}], Returns: Vector3
*/
func (o *Curve3D) GetPointPosition(idx gdnative.Int) gdnative.Vector3 {
	//log.Println("Calling Curve3D.GetPointPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_point_position")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns the tilt angle in radians for the point "idx". If the index is out of bounds, the function sends an error to the console, and returns 0.
	Args: [{ false idx int}], Returns: float
*/
func (o *Curve3D) GetPointTilt(idx gdnative.Int) gdnative.Float {
	//log.Println("Calling Curve3D.GetPointTilt()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_point_tilt")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Returns the position between the vertex "idx" and the vertex "idx"+1, where "t" controls if the point is the first vertex (t = 0.0), the last vertex (t = 1.0), or in between. Values of "t" outside the range (0.0 >= t <=1) give strange, but predictable results. If "idx" is out of bounds it is truncated to the first or last vertex, and "t" is ignored. If the curve has no points, the function sends an error to the console, and returns (0, 0, 0).
	Args: [{ false idx int} { false t float}], Returns: Vector3
*/
func (o *Curve3D) Interpolate(idx gdnative.Int, t gdnative.Float) gdnative.Vector3 {
	//log.Println("Calling Curve3D.Interpolate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromFloat(t)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "interpolate")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns a point within the curve at position "offset", where "offset" is measured as a distance in 3D units along the curve. To do that, it finds the two cached points where the "offset" lies between, then interpolates the values. This interpolation is cubic if "cubic" is set to true, or linear if set to false. Cubic interpolation tends to follow the curves better, but linear is faster (and often, precise enough).
	Args: [{ false offset float} {False true cubic bool}], Returns: Vector3
*/
func (o *Curve3D) InterpolateBaked(offset gdnative.Float, cubic gdnative.Bool) gdnative.Vector3 {
	//log.Println("Calling Curve3D.InterpolateBaked()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromFloat(offset)
	ptrArguments[1] = gdnative.NewPointerFromBool(cubic)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "interpolate_baked")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns the position at the vertex "fofs". It calls [method interpolate] using the integer part of fofs as "idx", and its fractional part as "t".
	Args: [{ false fofs float}], Returns: Vector3
*/
func (o *Curve3D) Interpolatef(fofs gdnative.Float) gdnative.Vector3 {
	//log.Println("Calling Curve3D.Interpolatef()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(fofs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "interpolatef")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Deletes the point "idx" from the curve. Sends an error to the console if "idx" is out of bounds.
	Args: [{ false idx int}], Returns: void
*/
func (o *Curve3D) RemovePoint(idx gdnative.Int) {
	//log.Println("Calling Curve3D.RemovePoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "remove_point")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false distance float}], Returns: void
*/
func (o *Curve3D) SetBakeInterval(distance gdnative.Float) {
	//log.Println("Calling Curve3D.SetBakeInterval()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(distance)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "set_bake_interval")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the position of the control point leading to the vertex "idx". If the index is out of bounds, the function sends an error to the console.
	Args: [{ false idx int} { false position Vector3}], Returns: void
*/
func (o *Curve3D) SetPointIn(idx gdnative.Int, position gdnative.Vector3) {
	//log.Println("Calling Curve3D.SetPointIn()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromVector3(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "set_point_in")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the position of the control point leading out of the vertex "idx". If the index is out of bounds, the function sends an error to the console.
	Args: [{ false idx int} { false position Vector3}], Returns: void
*/
func (o *Curve3D) SetPointOut(idx gdnative.Int, position gdnative.Vector3) {
	//log.Println("Calling Curve3D.SetPointOut()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromVector3(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "set_point_out")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the position for the vertex "idx". If the index is out of bounds, the function sends an error to the console.
	Args: [{ false idx int} { false position Vector3}], Returns: void
*/
func (o *Curve3D) SetPointPosition(idx gdnative.Int, position gdnative.Vector3) {
	//log.Println("Calling Curve3D.SetPointPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromVector3(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "set_point_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the tilt angle in radians for the point "idx". If the index is out of bounds, the function sends an error to the console. The tilt controls the rotation along the look-at axis an object traveling the path would have. In the case of a curve controlling a [PathFollow], this tilt is an offset over the natural tilt the PathFollow calculates.
	Args: [{ false idx int} { false tilt float}], Returns: void
*/
func (o *Curve3D) SetPointTilt(idx gdnative.Int, tilt gdnative.Float) {
	//log.Println("Calling Curve3D.SetPointTilt()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromFloat(tilt)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "set_point_tilt")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns a list of points along the curve, with a curvature controlled point density. That is, the curvier parts will have more points than the straighter parts. This approximation makes straight segments between each point, then subdivides those segments until the resulting shape is similar enough. "max_stages" controls how many subdivisions a curve segment may face before it is considered approximate enough. Each subdivision splits the segment in half, so the default 5 stages may mean up to 32 subdivisions per curve segment. Increase with care! "tolerance_degrees" controls how many degrees the midpoint of a segment may deviate from the real curve, before the segment has to be subdivided.
	Args: [{5 true max_stages int} {4 true tolerance_degrees float}], Returns: PoolVector3Array
*/
func (o *Curve3D) Tessellate(maxStages gdnative.Int, toleranceDegrees gdnative.Float) gdnative.PoolVector3Array {
	//log.Println("Calling Curve3D.Tessellate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(maxStages)
	ptrArguments[1] = gdnative.NewPointerFromFloat(toleranceDegrees)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "tessellate")

	// Call the parent method.
	// PoolVector3Array
	retPtr := gdnative.NewEmptyPoolVector3Array()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolVector3ArrayFromPointer(retPtr)
	return ret
}

// Curve3DImplementer is an interface that implements the methods
// of the Curve3D class.
type Curve3DImplementer interface {
	ResourceImplementer
	X_GetData() gdnative.Dictionary
	X_SetData(arg0 gdnative.Dictionary)
	AddPoint(position gdnative.Vector3, in gdnative.Vector3, out gdnative.Vector3, atPosition gdnative.Int)
	ClearPoints()
	GetBakeInterval() gdnative.Float
	GetBakedLength() gdnative.Float
	GetBakedPoints() gdnative.PoolVector3Array
	GetBakedTilts() gdnative.PoolRealArray
	GetPointCount() gdnative.Int
	GetPointIn(idx gdnative.Int) gdnative.Vector3
	GetPointOut(idx gdnative.Int) gdnative.Vector3
	GetPointPosition(idx gdnative.Int) gdnative.Vector3
	GetPointTilt(idx gdnative.Int) gdnative.Float
	Interpolate(idx gdnative.Int, t gdnative.Float) gdnative.Vector3
	InterpolateBaked(offset gdnative.Float, cubic gdnative.Bool) gdnative.Vector3
	Interpolatef(fofs gdnative.Float) gdnative.Vector3
	RemovePoint(idx gdnative.Int)
	SetBakeInterval(distance gdnative.Float)
	SetPointIn(idx gdnative.Int, position gdnative.Vector3)
	SetPointOut(idx gdnative.Int, position gdnative.Vector3)
	SetPointPosition(idx gdnative.Int, position gdnative.Vector3)
	SetPointTilt(idx gdnative.Int, tilt gdnative.Float)
	Tessellate(maxStages gdnative.Int, toleranceDegrees gdnative.Float) gdnative.PoolVector3Array
}