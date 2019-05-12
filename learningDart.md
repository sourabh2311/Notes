# Learning Dart!

* ```dart
  // comments
  /* */ multiline
  /** **/ or /// document comments
  ```

* A final variable can be set only once; a const variable is a compile-time constant. (Const variables are implicitly final.) A final top-level or class variable is initialized the first time it’s used.

* Use an assert statement to disrupt normal execution if a boolean condition is false.
  Note: Assert statements have no effect in production code; they’re for development only. Flutter enables asserts in debug mode.
  To attach a message to an assert, add a string as the second argument.

* Everything you can place in a variable is an object, and every object is an instance of a class. Even numbers, functions, and null are objects. Every object is an instance of a class, and all classes descend from Object (the base class for all Dart objects). So its same as in python. So int etc are immutable and list etc are mutable, so we have to be careful with that assignment (=) operator.

* You can use either single or double quotes to create a string

* var name = 'Bob';
  Variables store references. The variable called name contains a reference to a String object with a value of “Bob”.

  The type of the name variable is inferred to be String. If an object isn’t restricted to a single type, specify the Object or dynamic type, following design guidelines.
  dynamic name = 'Bob'; // so by adding dynamic we get python like functionality
  Another option is to explicitly declare the type that would be inferred:
  String name = 'Bob';

* Uninitialized variables have an initial value of null. Even variables with numeric types are initially null

* ```dart
  // String -> int
  var one = int.parse('1');
  assert(one == 1);
  ```

* ```dart
  // String -> double
  var onePointOne = double.parse('1.1');
  assert(onePointOne == 1.1);
  ```

* ```dart
  // int -> String
  String oneAsString = 1.toString();
  assert(oneAsString == '1');
  ```
* ```dart
  // double -> String
  String piAsString = 3.14159.toStringAsFixed(2); // leaving it empty will convert the whole to string
  assert(piAsString == '3.14');
  ```

* You can put the value of an expression inside a string by using ${expression}. If the expression is an identifier, you can skip the {}. To get the string corresponding to an object, Dart calls the object’s toString() method.

* ```dart
  var s1 = 'String '
    'concatenation'
    " works even over line breaks.";
  ```

* Another way to create a multi-line string: use a triple quote with either single or double quotation marks.

* You can create a “raw” string by prefixing it with r:
  `var s = r'In a raw string, not even \n gets special treatment.';`

* list is simply an array of same type objects with syntax similar to that of python

* set is same as in python
  ```dart
  var halogens = {'fluorine', 'chlorine', 'bromine', 'iodine', 'astatine'};
  // To create an empty set, use {} preceded by a type argument, or assign {} to a variable of type Set:
  var names = <String>{};
  // Set<String> names = {}; // This works, too.
  // var names = {}; // Creates a map, not a set as in python however in python we could create empty set like "x = set()" whose equivalent in dart is "var x = Set()".
  ```
* map is same as dictionary in python (in terms of syntax), Each key occurs only once, but you can use the same value multiple times. Both keys and values can be any type of object. Adding an object of different than Map<type of Key, type of Value> will give an error. Example:
  ```dart
  var gifts = {
    // Key:    Value
    'first': 'partridge',
    'second': 'turtledoves',
    'fifth': 'golden rings'
  };
  // aliter
  var gifts = Map(); // keyword "new" is optional
  gifts['first'] = 'partridge';
  gifts['second'] = 'turtledoves';
  gifts['fifth'] = 'golden rings';
  ```
* If you look for a key that isn’t in a map, you get a null in return.

* Use .length to get the number of key-value pairs in the map or the number of elements in a set.

* The usual way to express a Unicode code point is \uXXXX, where XXXX is a 4-digit hexadecimal value. For example, the heart character (♥) is \u2665. To specify more or less than 4 hex digits, place the value in curly brackets. For example, the laughing emoji (😆) is \u{1f600}.

* ```dart
  isNoble(atomicNumber) { // we could have added "bool" on the left of "isNobel (..) {"
    return _nobleGases[atomicNumber] != null;
  }
  ```

* All functions return a value. If no return value is specified, the statement return null; is implicitly appended to the function body.

* For functions that contain just one expression, you can use a shorthand syntax:

  `bool isNoble(int atomicNumber) => _nobleGases[atomicNumber] != null;`

  The => expr syntax is a shorthand for { return expr; }. 

  Note: Only an expression—not a statement—can appear between the arrow (=>) and the semicolon (;). For example, you can’t put an if statement there, but you can use a conditional expression.

* A function can have two types of parameters: required and optional. The required parameters are listed first, followed by any optional parameters. Optional parameters can be either positional or named, but not both.

* **Optional named parameters** When defining a function, use `{param1, param2, ...}` to specify named parameters. When calling a function, you can specify named parameters using paramName: value. Example:
  ```dart
  enableFlags(bold: true, hidden: false); 
  /// Sets the [bold] and [hidden] flags in ...
  void enableFlags({bool bold, bool hidden}) {...}
  ```

  You can annotate a named parameter in any Dart code (not just Flutter) with @required to indicate that it is a required parameter.

  `const Scrollbar({Key key, @required Widget child})`

  When a Scrollbar is constructed, the analyzer reports an issue when the child argument is absent. 

* Wrapping a set of function parameters in [] marks them as optional positional parameters:

  ```dart
  String say(String from, String msg, [String device]) {
    var result = '$from says $msg';
    if (device != null) {
      result = '$result with a $device';
    }
    return result;
  }
  ```
* Default parameter values: Ex: 

  ```dart
  // Sets the [bold] and [hidden] flags ...
  void enableFlags({bool bold = false, bool hidden = false}) {...}
  // bold will be true; hidden will be false.
  enableFlags(bold: true); // writing just true wont work i.e. we must specify the name
  ```

* If no default value is provided, the default value is null.

* The next example shows how to set default values for positional parameters:

  ```dart
  String say(String from, String msg, [String device = 'carrier pigeon', String mood]) {
    //...
  }
  ```

* Most functions are named, such as main() or printElement(). You can also create a nameless function called an anonymous function, or sometimes a lambda or closure.
  They have the format:
  ```dart
  ([[Type] param1[, …]]) {
    codeBlock;
  };
  // example:
  var list = ['apples', 'bananas', 'oranges'];
  list.forEach((item) {
    print('${list.indexOf(item)}: $item');
  });
  Output:
  0: apples
  1: bananas
  2: oranges
  ```
* ```dart
  /// Returns a function that adds [addBy] to the
  /// function's argument.
  Function makeAdder(num addBy) {
    return (num i) => addBy + i; // A function expression cant have a name
  }
  /* This is also valid
  makeAdder(num addBy) {
    return (num i) { // couldnt have done return int name (num i) or even return name (num i)
      return addBy + i;
    };
  }
  */

* /   Divide (double result)

* ~/  Divide (integer result)

* %   Get the remainder of an integer division (modulo)

* Warning: For operators that work on two operands, the leftmost operand determines which version of the operator is used. For example, if you have a Vector object and a Point object, aVector + aPoint uses the Vector version of +.
* Type Test operators (checking types at runtime):

  * `as`  Typecast

  * `is`  True if the object has the specified type

  * `is!` False if the object has the specified type

* `expr1 ?? expr2` If expr1 is non-null, returns its value; otherwise, evaluates and returns the value of expr2.
  ```dart
  // Assign value to b if b is null; otherwise, b stays the same
  b ??= value;
  // one more example
  String playerName(String name) => name ?? 'Guest';
  ``` 
* `?.`  'Conditional member access' - Like ., but the leftmost operand can be null; example: foo?.bar selects property bar from expression foo unless foo is null (in which case the value of foo?.bar is null)

* Iterable classes such as List and Set also support the for-in form of iteration:
  ```dart
  var collection = [0, 1, 2];
  for (var x in collection) {
    print(x); // 0 1 2
  }
  ```

* for, while, if, break, continue, do while, switch case are same as in c

* Exceptions work similar to other languages:
  ```dart
  // throw an exception like 
  throw myExceptionName('Some String');
  // we can also throw arbitrary objects
  throw 'Some String';
  // try and catch
  try {
    breedMoreLlamas();
  } on OutOfLlamasException {
    // A specific exception
    buyMoreLlamas();
  } on Exception catch (e) {
    // Anything else that is an exception
    print('Unknown exception: $e');
  } catch (e) {
    // No specified type, handles all
    print('Something really unknown: $e');
  }
  ```

  Note: The first catch clause that matches the thrown object’s type handles the exception.
  You can specify one or two parameters to catch(). The first is the exception that was thrown, and the second is the stack trace (a StackTrace object).
* To partially handle an exception, while allowing it to propagate, use the rethrow keyword.

  ```dart
  void misbehave() {
    try {
      dynamic foo = true;
      print(foo++); // Runtime error
    } catch (e) {
      print('misbehave() partially handled ${e.runtimeType}.'); // runtimeType gives the type of an object at runtime.
      rethrow; // Allow callers to see the exception.
    }
  }
  void main() {
    try {
      misbehave();
    } catch (e) {
      print('main() finished handling ${e.runtimeType}.');
    }
  }
  ```

* To ensure that some code runs whether or not an exception is thrown, use a finally clause. If no catch clause matches the exception, the exception is propagated after the finally clause runs.

* The finally clause runs after any matching catch clauses

* You can create an object using a constructor. Constructor names can be either ClassName or ClassName.identifier.
  ```dart
  var p1 = Point(2, 2); // we can optionally add the 'new' keyword before Point
  var p2 = Point.fromJson({'x': 1, 'y': 2}); // Map<String, int>
  ``` 
* Some classes provide constant constructors. To create a compile-time constant using a constant constructor, put the const keyword before the constructor name:
  ```dart
  var p = const ImmutablePoint(2, 2);
  ```

* Within a constant context, you can omit the const before a constructor or literal. For example, look at this code, which creates a const map:
  ```dart
  // Lots of const keywords here.
  const pointAndLine = const {
    'point': const [const ImmutablePoint(0, 0)],
    'line': const [const ImmutablePoint(1, 10), const ImmutablePoint(-2, 11)],
  };

  // You can omit all but the first use of the const keyword:
  // Only one const, which establishes the constant context.
  const pointAndLine = {
    'point': [ImmutablePoint(0, 0)],
    'line': [ImmutablePoint(1, 10), ImmutablePoint(-2, 11)],
  };
  ```

* In class, all uninitialized instance variables have the value null and all instance variables generate an implicit getter method. Non-final instance variables also generate an implicit setter method.
  ```dart
  // example to understand getter and setter.
  class Rectangle {
    num left, top, width, height;

    Rectangle(this.left, this.top, this.width, this.height);

    // Define two calculated properties: right and bottom.
    num get right => left + width;
    set right(num value) => left = value - width;
    num get bottom => top + height;
    set bottom(num value) => top = value - height;
  }

  void main() {
    var rect = Rectangle(3, 4, 20, 15);
    assert(rect.left == 3);
    rect.right = 12;
    assert(rect.left == -8);
  }

* ```dart
  // Constructors
  class Point {
    num x, y;

    Point(num x, num y) {
      // There's a better way to do this, stay tuned.
      this.x = x;
      this.y = y;
    }
  }


  // The pattern of assigning a constructor argument to an instance variable is so common, Dart has syntactic sugar to make it easy:

  class Point {
    num x, y;
    // Syntactic sugar for setting x and y
    // before the constructor body runs.
    Point(this.x, this.y);
    // use named constructor to implement multiple constructors for a class
    Point.origin() { x = 0; y = 0; }
  }
  ```

* **Initializer list** Besides invoking a superclass constructor, you can also initialize instance variables before the constructor body runs. Separate initializers with commas. They are as well handy in setting up final fields. And during development, you can validate inputs by using assert in the initializer list.

  ```dart
  // Initializer list sets instance variables before
  // the constructor body runs.
  Point.fromJson(Map<String, num> json)
      : x = json['x'],
        y = json['y'] {
    print('In Point.fromJson(): ($x, $y)');
  }
  ```
* If you don’t declare a constructor, a default constructor is provided for you. The default constructor has no arguments and invokes the no-argument constructor in the superclass.

* Constructors aren’t inherited: Subclasses don’t inherit constructors from their superclass. A subclass **that declares no constructors** (so if you are declaring a constructor don't have a default constructor) has only the default (no argument, no name) constructor. 

* By default, a constructor *(any\*)* in a subclass calls the superclass’s unnamed, no-argument constructor. The superclass’s constructor is called at the beginning of the constructor body. If an initializer list is also being used, it executes before the superclass is called. In summary, the order of execution is as follows:
  1.    initializer list
  2.    superclass’s no-arg constructor
  3.    main class’s no-arg constructor

  If the superclass doesn’t have an unnamed, no-argument constructor, then you must manually call one of the constructors in the superclass. Specify the superclass constructor after a colon (:), just before the constructor body (if any).
  ```dart
  class Person {
    String firstName;

    Person.fromJson(Map data) {
      print('in Person');
    }
  }

  class Employee extends Person {
    // Person does not have a default constructor;
    // you must call super.fromJson(data).
    Employee.fromJson(Map data) : super.fromJson(data) {
      print('in Employee');
    }
  }
  main() {
    var emp = new Employee.fromJson({});

    // Prints:
    // in Person
    // in Employee
    if (emp is Person) {
      // Type check
      emp.firstName = 'Bob';
    }
    (emp as Person).firstName = 'Bob';
  }
  ```
* Sometimes a constructor’s only purpose is to redirect to another constructor in the same class. A redirecting constructor’s body is empty, with the constructor call appearing after a colon (:).
  ```dart
  class Point {
    num x, y;

    // The main constructor for this class.
    Point(this.x, this.y);

    // Delegates to the main constructor.
    Point.alongXAxis(num x) : this(x, 0);
  }
  ```

* If your class produces objects that never change, you can make these objects compile-time constants. To do this, define a const constructor and make sure that all instance variables are final.
  ```dart
  class ImmutablePoint {
    static final ImmutablePoint origin =
        const ImmutablePoint(0, 0);

    final num x, y;

    const ImmutablePoint(this.x, this.y);
  }
  // Remember:
  var a = const ImmutablePoint(1, 1); // Creates a constant
  var b = ImmutablePoint(1, 1); // Does NOT create a constant
  ```
* Use the factory keyword when implementing a constructor that doesn’t always create a new instance of its class. For example, a factory constructor might return an instance from a cache, or it might return an instance of a subtype.

  The following example demonstrates a factory constructor returning objects from a cache:
  ```dart
  class Logger {
    final String name;
    bool mute = false;

    // _cache is library-private, thanks to
    // the _ in front of its name.
    static final Map<String, Logger> _cache =
        <String, Logger>{};

    factory Logger(String name) {
      if (_cache.containsKey(name)) {
        return _cache[name];
      } else {
        final logger = Logger._internal(name);
        _cache[name] = logger;
        return logger;
      }
    }

    Logger._internal(this.name);

    void log(String msg) {
      if (!mute) print(msg);
    }
  }
  // Note: Factory constructors have no access to this.
  ```
  Invoke a factory constructor just like you would any other constructor:
  ```dart
  var logger = Logger('UI');
  logger.log('Button clicked');
  ```
* Instance, getter, and setter methods can be abstract, defining an interface but leaving its implementation up to other classes. Abstract methods can only exist in abstract classes.

  To make a method abstract, use a semicolon (;) instead of a method body:
  ```dart
  abstract class Doer {
    // Define instance variables and methods...

    void doSomething(); // Define an abstract method.
  }

  class EffectiveDoer extends Doer {
    void doSomething() {
      // Provide an implementation, so the method is not abstract here...
    }
  } 
  ```
* An **abstract class**—a class that can’t be instantiated. Abstract classes are useful for defining interfaces, often with some implementation. If you want your abstract class to appear to be instantiable, define a factory constructor.
* Every class implicitly defines an interface containing all the instance members of the class and of any interfaces it implements. If you want to create a class A that supports class B’s API without inheriting B’s implementation, class A should implement the B interface.

  A class implements one or more interfaces by declaring them in an implements clause and then providing the APIs required by the interfaces. For example:
  ```dart
  // A person. The implicit interface contains greet().
  class Person {
    // In the interface, but visible only in this library.
    final _name;

    // Not in the interface, since this is a constructor.
    Person(this._name);

    // In the interface.
    String greet(String who) => 'Hello, $who. I am $_name.';
  }

  // An implementation of the Person interface.
  class Impostor implements Person {
    get _name => '';

    String greet(String who) => 'Hi $who. Do you know who I am?';
  }

  String greetBob(Person person) => person.greet('Bob');

  void main() {
    print(greetBob(Person('Kathy')));
    print(greetBob(Impostor()));
  }
  ```
  Here’s an example of specifying that a class implements multiple interfaces:
  ```dart
  class Point implements Comparable, Location {...}
  ```
* Use extends to create a subclass, and super to refer to the superclass.
  ```dart
  class Television {
    void turnOn() {
      _illuminateDisplay();
      _activateIrSensor();
    }
    // ···
  }

  class SmartTelevision extends Television {
    void turnOn() {
      super.turnOn();
      _bootNetworkInterface();
      _initializeMemory();
      _upgradeApps();
    }
    // ···
  }
  ```
* Subclasses can override instance methods, getters, and setters. You can use the @override annotation to indicate that you are intentionally overriding a member. We can even override the binary operators.
  ```dart
  class SmartTelevision extends Television {
    @override  // so as we saw in the above example this is optional. Even when overriding we can access supers turnOn.
    void turnOn() {...}
    // ···
  }
  ```

  Operator overloading:
  ```dart
  class Vector {
    final int x, y;

    Vector(this.x, this.y);

    Vector operator +(Vector v) => Vector(x + v.x, y + v.y);
    Vector operator -(Vector v) => Vector(x - v.x, y - v.y);

    // Operator == and hashCode not shown. For details, see note below.
    // ···
  }
  ```
  if you override ==, you should also override Object’s hashCode getter.

* To detect or react whenever code attempts to use a non-existent method or instance variable, you can override noSuchMethod():
  ```dart
  class A {
    // Unless you override noSuchMethod, using a
    // non-existent member results in a NoSuchMethodError.
    @override
    void noSuchMethod(Invocation invocation) {
      print('You tried to use a non-existent member: ' +
          '${invocation.memberName}');
    }
  }
  ```
* Use the static keyword to implement class-wide variables and methods.

  Static variables (class variables) are useful for class-wide state and constants:
  ```dart
  class Queue {
    static const initialCapacity = 16;
    // ···
  }

  void main() {
    assert(Queue.initialCapacity == 16);
  }
  ```
  Static variables aren’t initialized until they’re used.
  
  Static methods (class methods) do not operate on an instance, and thus do not have access to this. For example:
  ```dart
  class Point {
    num x, y;
    Point(this.x, this.y);

    static num distanceBetween(Point a, Point b) {
      var dx = a.x - b.x;
      var dy = a.y - b.y;
      return sqrt(dx * dx + dy * dy);
    }
  }

  void main() {
    var a = Point(2, 2);
    var b = Point(4, 4);
    var distance = Point.distanceBetween(a, b);
    assert(2.8 < distance && distance < 2.9);
    print(distance);
  }
  ```
* Generics are same as templates, example:
  ```dart
  abstract class Cache<T> {
    T getByKey(String key);
    void setByKey(String key, T value);
  }
  ```
* When implementing a generic type, you might want to limit the types of its parameters. You can do this using extends.
  ```dart
  class Foo<T extends SomeBaseClass> {
    // Implementation goes here...
    String toString() => "Instance of 'Foo<$T>'";
  }

  class Extender extends SomeBaseClass {...}
  ```
  It’s OK to use SomeBaseClass or any of its subclasses as generic argument:
  ```dart
  var someBaseClassFoo = Foo<SomeBaseClass>();
  var extenderFoo = Foo<Extender>();
  ```
  It’s also OK to specify no generic argument:
  ```dart
  var foo = Foo();
  print(foo); // Instance of 'Foo<SomeBaseClass>'
  ```
  Specifying any non-SomeBaseClass type results in an error:
  ```dart
  var foo = Foo<Object>();
  ```
* Generic methods
  ```dart
  T first<T>(List<T> ts) {
    // Do some initial work or error checking, then...
    T tmp = ts[0];
    // Do some additional checking or processing...
    return tmp;
  }
  ```
* Libraries not only provide APIs, but are a unit of privacy: identifiers that start with an underscore (_) are visible only inside the library. Every Dart app is a library, even if it doesn’t use a library directive.

* If you import two libraries that have conflicting identifiers, then you can specify a prefix for one or both libraries.

  ```dart
  import 'package:lib1/lib1.dart';
  import 'package:lib2/lib2.dart' as lib2;

  // Uses Element from lib1.
  Element element1 = Element();

  // Uses Element from lib2.
  lib2.Element element2 = lib2.Element();
  ```

* If you want to use only part of a library, you can selectively import the library.
  ```dart
  // Import only foo.
  import 'package:lib1/lib1.dart' show foo;

  // Import all names EXCEPT foo.
  import 'package:lib2/lib2.dart' hide foo;
  ```

* To lazily load a library, you must first import it using deferred as.
  ```dart
  import 'package:greetings/hello.dart' deferred as hello;

  // When you need the library, invoke loadLibrary() using the library’s identifier.

  Future greet() async {
    await hello.loadLibrary();
    hello.printGreeting();
  }
  ```
  In the preceding code, the await keyword pauses execution until the library is loaded.
  You can invoke loadLibrary() multiple times on a library without problems. The library is loaded only once.
* 
  ```dart
  import 'dart:async';


  main() {
    print("Main starts");
    printFileContent();
    print("Main ends");
  }

  printFileContent() async { // we must add async if we are using await
    String fileContent = await downloadFile(); // will wait until the file is ready
    print("FC: $fileContent");
  }
  /* 
  // Following is equivalent
  printFileContent() {
    Future<String> fileContent = downloadFile();
    fileContent.then((resultString) {
      print("FC: $fileContent");
    })
  }
  */

  Future<String> downloadFile() {
    return Future.delayed(Duration(seconds: 6), () {
      return "My FC";
    });
  }
  /* output:
  Main starts
  Main ends
  FC: My FC // got printed after 6 seconds
  */
  ```