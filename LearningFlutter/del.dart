import 'dart:async';


main() {
  print("Main starts");
  printFileContent();
  print("Main ends");
}

printFileContent() async { // we must add async if we are using await
  String fileContent = await downloadFile();
  print("FC: $fileContent");
}

Future<String> downloadFile() {
  return Future.delayed(Duration(seconds: 6), () {
    return "My FC";
  });
}