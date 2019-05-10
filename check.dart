class television {
  int x;
  television (this.x);
  void turnon() {
    print(x);
  }
}

class smartTelevision extends television {
  smartTelevision (int x) : super(x);
  void turnon() {
    super.turnon();
    print("I am smart");
  } 
}

class smartTelevisionV2 extends television {
  smartTelevisionV2 (int x) : super(x);
  @override
  void turnon() {
    super.turnon();
    print("I am neosmart");
  } 
}

void main () {
  var mytv = smartTelevision(3);
  mytv.turnon();
  var neotv = smartTelevisionV2(4);
  neotv.turnon();
}