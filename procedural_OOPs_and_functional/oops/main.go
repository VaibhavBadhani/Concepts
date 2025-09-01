package main

import "fmt"

// Define a struct to represent the data (fields) of a Rectangle object.
type Rectangle struct {
    length float64
    width  float64
}

// Attach a method to the Rectangle struct to calculate its area.
// The (r Rectangle) part is the receiver, which makes this a method of the struct.
func (r Rectangle) Area() float64 {
    return r.length * r.width
}

// Attach another method to calculate the perimeter.
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.length + r.width)
}

func main() {
    // 1. Create an instance of the Rectangle object.
    myRectangle := Rectangle{length: 10.5, width: 5.2}

    // 2. Call the methods on the object to perform actions.
    area := myRectangle.Area()
    perimeter := myRectangle.Perimeter()

    // 3. Print the results.
    fmt.Printf("The area of the rectangle is: %.2f\n", area)
    fmt.Printf("The perimeter of the rectangle is: %.2f\n", perimeter)
}









// 1. Encapsulation -------------------------------------------------------------------

// Encapsulation is about bundling data with the methods that operate on that data.
// A single node in an SLL is a great example. It combines its data (value) and a 
// pointer to the next node (next) into a single, self-contained unit.



// Node encapsulates its data and a pointer to the next node.
type Node struct {
    value int
    next  *Node
}

// In this code, Node encapsulates the value and the next pointer, keeping them 
// together as a single entity. The user of the Node doesn't need to know how these
//  are stored; they just interact with the Node as a whole.

// 2. Abstraction -------------------------------------------------------------------

// Abstraction involves hiding complex implementation details and showing only the 
// essential features of an object. For an SLL, a user doesn't need to know the 
// intricate memory management of pointers. They only need a simple set of operations
//  like add, remove, or find.

// LinkedList provides an abstract interface for the user.
type LinkedList struct {
    head *Node
}

// Add is an abstract operation for the user. They don't need to know
// how the new node is linked to the rest of the list.
func (l *LinkedList) Add(value int) {
    newNode := &Node{value: value}
    if l.head == nil {
        l.head = newNode
        return
    }
    current := l.head
    for current.next != nil {
        current = current.next
    }
    current.next = newNode
}

// The user of the LinkedList only needs to call Add(value). The underlying logic of
//  traversing the list and setting pointers is abstracted away, making the list easy
//  to use.

// 3. Inheritance -------------------------------------------------------------------

// Inheritance is when a new class (subclass) inherits properties and methods from an
//  existing class (parent class). SLLs, as a simple data structure, don't typically
//  use inheritance on their own. However, you can use Go's struct embedding to achieve
//  a form of inheritance. For example, a DoublyLinkedList could be a type that embeds
//  a LinkedList to inherit its basic properties, like the head pointer.

// This is not a classic SLL, but shows how a similar concept can be extended
// using Go's struct embedding, which is a form of composition.

type DoubleNode struct {
    Node // This is struct embedding, providing inheritance-like behavior.
    prev *DoubleNode
}

// Here, DoubleNode "inherits" the value and next fields from the Node struct, plus it 
// adds its own prev field.

// 4. Polymorphism -------------------------------------------------------------------

// Polymorphism means "many forms," and it allows objects to be treated as instances of
//  their parent type. In Go, this is primarily achieved through interfaces. A function
//  can accept a parameter of an interface type, and it will work with any object that
//  satisfies that interface, regardless of its underlying struct type.

// IPrintable is an interface that can be implemented by different structs.
type IPrintable interface {
    Print()
}

// Node can satisfy the IPrintable interface.
func (n *Node) Print() {
    fmt.Printf("Node value: %d\n", n.value)
}

// AnotherStruct can also satisfy the IPrintable interface.
type AnotherStruct struct {
    data string
}

func (a *AnotherStruct) Print() {
    fmt.Println("Another struct data:", a.data)
}

// This function is polymorphic, as it can accept any object that
// satisfies the IPrintable interface.
func printObject(p IPrintable) {
    p.Print()
}

// The printObject function can accept a Node or an AnotherStruct because both implement
//  the IPrintable interface. When printObject is called, it will execute the correct 

// Print() method for the specific object it receives. This is a powerful demonstration
//  of polymorphism.