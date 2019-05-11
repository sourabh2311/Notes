# Learning Flutter!

* Everything in Flutter is a Widget and a Widget is nothing more than a Dart class that extends a Flutter class.
* Widget classes have (usually) only one requirement: it must have a build method which returns other Widgets. The only exception to this rule is low-level widgets like 'Text' that return primitive types (Strings or numbers, usually.)

  Other than that, a widget is just a normal Dart class. You can add methods and properties and the like.
  ```dart
  class BigText extends StatelessWidget {
    // a property on this class
    final String text;

    // a constructor for this class
    BigText(this.text);

    Widget build(context) {
      // Pass the text down to another widget
      return new Text(
        text,
        // Even changing font-style is done through a Dart class.
        textStyle: new TextStyle(fontSize: 20.0),
      );
    }
  }
  // Then somewhere else in an app you'd use the widget like this:

  // ...
  // This is how we'd use the BigText within another widget.
  child: new BigText('This string would render and be big'),
  // ...
  ```
* Material widgets are designed to look like Android apps and Cupertino like iOS.
* Flutter widgets must extend a handful of classes from the Flutter library. The two you'll use almost always are StatelessWidget and StatefulWidget.

  The difference is that one has a concept of state within the Widget and some built in methods that tells Flutter to re-render if that state changes.
  
  A Stateful Widget looks a bit different. It's actually two classes: the state object and the widget.

## Lesson 1

* Wrote these two files:
* main.dart
  ```dart
  import "package:flutter/material.dart";
  import './app_screens/first_screen.dart';

  void main() => runApp(new MyFlutterApp());

  class MyFlutterApp extends StatelessWidget {

    @override
    Widget build(BuildContext context) {
      return MaterialApp(
          debugShowCheckedModeBanner: false,
          title: "My Flutter App",
          home: Scaffold(
            appBar: AppBar(title: Text("My First App Screen"),),
            body: FirstScreen()
          )
      );
    }
  }
  ```
  And app_screens/first_screen.dart
  ```dart
  import 'dart:math';

  import 'package:flutter/material.dart';

  class FirstScreen extends StatelessWidget {

    @override
    Widget build(BuildContext context) {
      return Material(
        color: Colors.lightBlueAccent,
        child: Center(
          child: Text(
            generateLuckyNumber(),
            textDirection: TextDirection.ltr,
            style: TextStyle(color: Colors.white, fontSize: 40.0),
          ),
        ),
      );
    }

    String generateLuckyNumber() {

      var random = Random();
      int luckyNumber = random.nextInt(10);
      return "Your lucky number is $luckyNumber";
    }
  }
  ```
## Lesson 2
* main.dart
  ```dart
  import 'package:flutter/material.dart';
  import './app_screens/home.dart';

  void main() {
    runApp(MaterialApp(
      title: "Exploring UI widgets",
      home: Home(),
    ));
  }
  ```
  The other file
  ```dart
  import 'package:flutter/material.dart';

  class Home extends StatelessWidget {
    @override
    Widget build(BuildContext context) {
      return Center(
          child: Container(
              // for working with fonts, create a new folder at application root named 'fonts' then put your downloaded fonts there and add their details in pubspec.yaml. Note: for regular font there is no need to define weight but for italic and bold find weight when downloading the font and mention their style as well (not need for regular font)
              // If we hadn't used Center widget then these width and height settings wouldn't work. But after using Center they would work, also margin is distance between two widgets and padding is like distance from boundary. Similary we have EdgeInsets.all() to get margin/padding from everywhere
    //		    width: 200.0,
    //		    height: 100.0,
    //		    margin: EdgeInsets.only(left: 35.0, top: 50.0),
              // added padding to get space for text. 
              padding: EdgeInsets.only(left: 10.0, top: 40.0),
              alignment: Alignment.center,
              color: Colors.deepPurple,
              child: Column(
                // Column is like |r1|
                //                |r2|
                //                |r3|
                // Whereas row is like |c1-c2-c3|
                // Our text was too big to get fit and therefore we started using expanded.
                children: <Widget>[
                  Row(
                    children: <Widget>[
                      Expanded(
                          child: Text(
                        "Spice Jet",
                        textDirection: TextDirection.ltr,
                        style: TextStyle(
                            decoration: TextDecoration.none,
                            fontSize: 35.0,
                            fontFamily: 'Raleway',
                            fontWeight: FontWeight.w700,
                            color: Colors.white),
                      )),
                      Expanded(
                          child: Text(
                        "From Mumbai to Bangalore via New Delhi",
                        textDirection: TextDirection.ltr,
                        style: TextStyle(
                            decoration: TextDecoration.none,
                            fontSize: 20.0,
                            fontFamily: 'Raleway',
                            fontWeight: FontWeight.w700,
                            color: Colors.white),
                      )),
                    ],
                  ),
                  Row(
                    children: <Widget>[
                      Expanded(
                          child: Text(
                        "Air India",
                        textDirection: TextDirection.ltr,
                        style: TextStyle(
                            decoration: TextDecoration.none,
                            fontSize: 35.0,
                            fontFamily: 'Raleway',
                            fontWeight: FontWeight.w700,
                            color: Colors.white),
                      )),
                      Expanded(
                          child: Text(
                        "From Jaipur to Goa",
                        textDirection: TextDirection.ltr,
                        style: TextStyle(
                            decoration: TextDecoration.none,
                            fontSize: 20.0,
                            fontFamily: 'Raleway',
                            fontWeight: FontWeight.w700,
                            color: Colors.white),
                      )),
                    ],
                  ),
                  FlightImageAsset(),
                  FlightBookButton()
                ],
              )));
    }
  }
  // Adding an image is simple! Just create folder images like fonts and put your images there and enter the details in that pubspec.yaml file.
  class FlightImageAsset extends StatelessWidget {

    @override
    Widget build(BuildContext context) {
      AssetImage assetImage = AssetImage('images/flight.png');
      Image image = Image(image: assetImage, width: 250.0, height: 250.0,);
      return Container(child: image,);
    }
  }
  // book button
  class FlightBookButton extends StatelessWidget {
    @override
    Widget build(BuildContext context) {
      return Container(
        // space from other widget
        margin: EdgeInsets.only(top: 30.0),
        // defining the containers size
        width: 250.0,
        height: 50.0,
        child: RaisedButton(
            color: Colors.deepOrange,
            child: Text(
              "Book Your Flight",
              style: TextStyle(
                  fontSize: 20.0,
                  color: Colors.white,
                  fontFamily: 'Raleway',
                  fontWeight: FontWeight.w700),
            ),
            // will make button look like its elevated from the container (putting the appropriate shadow)
            elevation: 6.0,
            onPressed: () => bookFlight(context)),
      );
    }

    void bookFlight(BuildContext context) {
      var alertDialog = AlertDialog(
        title: Text("Flight Booked Successfully"),
        content: Text("Have a pleasant flight"),
      );

      showDialog(
          context: context,
          builder: (BuildContext context) => alertDialog);
    }
  }
  ```
  Output:

  <img src="image2.png" alt="drawing" width="200"/>
  <img src="image3.png" alt="drawing" width="200"/>
* Basic List View. Notes:-
  * Meant for small number of items as it loads all the items in memory when attached to the screen.
  * Always wrap ListView as 'home' of 'Scaffold' widget as it is scrollable and might overflow beyond the screen 
  * Example code:
  ```dart
  void main() {
    runApp(MaterialApp(
      title: "Exploring UI widgets",
      home: Scaffold(
        appBar: AppBar(title: Text("Basic List View"),),
        body: getListView(),
      ),
    ));
  }

  Widget getListView() {

    var listView = ListView(
      children: <Widget>[

        ListTile(
          leading: Icon(Icons.landscape),
          title: Text("Landscape"),
          subtitle: Text("Beautiful View !"),
          trailing: Icon(Icons.wb_sunny),
          onTap: () {
            debugPrint("Landscape tapped"); // will just print on terminal
          },
        ),

        ListTile(
          leading: Icon(Icons.laptop_chromebook),
          title: Text("Windows"),
        ),

        ListTile(
          leading: Icon(Icons.phone),
          title: Text("Phone"),
        ),

        Text("Yet another element in List"),

        Container(color: Colors.red, height: 50.0,)

      ],
    );

    return listView;
  }

  ```
  Output:

  <img src="image1.png" alt="drawing" width="200"/>

* Implementing long lists
  ```dart
  void main() {
    runApp(MaterialApp(

      title: "Exploring UI widgets",

      home: Scaffold(
        appBar: AppBar(title: Text("Long List"),),
        body: getListView(),
        // Remember: 1 - Every Widget has its own build() and its context. 2 - The BuildContext is the parent of the widget returned by the build() method.
        // FAB is available only inside Scaffolds
        floatingActionButton: FloatingActionButton(
          onPressed: () {
            debugPrint("FAB clicked");
          },
          child: Icon(Icons.add),
          tooltip: 'Add One More Item', // will show after a long pressed on that button
        ),
      ),

    ));
  }

  void showSnackBar(BuildContext context, String item) {
    var snackBar = SnackBar(
      content: Text("You just tapped $item"),
      action: SnackBarAction(
          label: "UNDO",
          onPressed: () {
            debugPrint('Performing dummy UNDO operation');
          }
      ),
    );
    // like FAB snackbar needs context of our Scaffold
    Scaffold.of(context).showSnackBar(snackBar);
  }

  // preparing our data source
  List<String> getListElements() {

    var items = List<String>.generate(1000, (counter) => "Item $counter");
    return items;
  }

  Widget getListView() {

    var listItems = getListElements();

    // this is memory efficient as list will only be built for those items that can fit on screen
    var listView = ListView.builder(
        itemBuilder: (context, index) {

          return ListTile(
            leading: Icon(Icons.arrow_right),
            title: Text(listItems[index]),
            onTap: () {
              showSnackBar(context, listItems[index]);
            },
          );
        }
    );

    return listView;
  }
  ```

  Output:

  <img src="image4.png" alt="drawing" width="200"/>
  <img src="image5.png" alt="drawing" width="200"/>