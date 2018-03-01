package inputevent

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
Contains a generic action which can be targeted from several type of inputs. Actions can be created from the project settings menu [code]Project > Project Settings > Input Map[/code]. See [method Node._input].
*/
type InputEventAction struct {
	InputEvent
}

func (o *InputEventAction) BaseClass() string {
	return "InputEventAction"
}

/*
   Undocumented
*/
func (o *InputEventAction) GetAction() gdnative.String {
	log.Println("Calling InputEventAction.GetAction()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_action", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *InputEventAction) SetAction(action gdnative.String) {
	log.Println("Calling InputEventAction.SetAction()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(action)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_action", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *InputEventAction) SetPressed(pressed gdnative.Bool) {
	log.Println("Calling InputEventAction.SetPressed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(pressed)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_pressed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   InputEventActionImplementer is an interface for InputEventAction objects.
*/
type InputEventActionImplementer interface {
	Class
}