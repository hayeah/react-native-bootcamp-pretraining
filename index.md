# Sike.io ReactNative Bootcamp Fundamentals

To get ready for our React/ReactNative bootcamp, you don't need to know web development very well, but you do need to know a little about a lot of things.

That's why we prepared this checklist to help you learn the fundamentals before we start the bootcamp.

This checklist is a list of questions. Take a quick look and see if you know the answers. If you see anything you are not sure about, read the learning resources we've provided. If everything in this checklist is familiar to you, then you are ready.

+ If you are new to web development, you should spend as much time as you can learning JavaScript.
+ If you are a frontend developer, you should spend your time learning iOS fundamentals.

# JavaScript

Can you understand this piece of code?

```
/**
* This is a function that calls the callback at a fixed interval.
*
* @param interval Time interval between each counter increment.
* @param callback Pass the counter value to the callback at every interval.
* @return A function to cancel the counter.
*/
function startCounter(interval,callback) {
  var i = 0;

  // Do you know what an anonymous function is?
  var timerID = setInterval(function() {
    i++;
    callback(i);
  },interval);


  // Do you understand how closure captures the timerID?
  var cancel = function() {
    clearInterval(timerID);
  };

  // Do you know that closure/function can be a value?
  return cancel;
}

// Do you know what a callback function is?
var cancel = startCounter(500,function(i) {
  console.log("currenet counter value:",i);
});

// stop the counter after 5 seconds.
setTimeout(cancel,5000);


// Output
// currenet counter value: 1
// currenet counter value: 2
// currenet counter value: 3
// currenet counter value: 4
// currenet counter value: 5
// currenet counter value: 6
// currenet counter value: 7
// currenet counter value: 8
// currenet counter value: 9

```

If you are not familiar with JavaScript syntax, read:

+ [快速入门](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/00143449917624134f5c4695b524e81a581ab5a222b05ec000)
+ [函数定义和调用](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/00143449926746982f181557d9b423f819e89709feabdb4000)

If you are not familiar with closure and anonymous function, read:

+ [闭包](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/00143449934543461c9d5dfeeb848f5b72bd012e1113d15000)
+ [变量作用域](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/0014344993159773a464f34e1724700a6d5dd9e235ceb7c000)

### Objects in JavaScript

Can you understand this piece of code?

```
// Person is a constructor
function Person(name) {
  this.name;
}

// Every instance of person would have the `sayHello` method.
Person.prototype.sayHello = function() {
  console.log("hello, my name is " + this.name);
}

// Create an instance of Person using the Peron constructor.
var person = new Person();

// Calls person's method.
person.sayHello();
```

+ [方法](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/0014345005399057070809cfaa347dfb7207900cfd116fb000)
+ [创建对象](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/0014344997235247b53be560ab041a7b10360a567422a78000)
+ [原型继承](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/0014344997013405abfb7f0e1904a04ba6898a384b1e925000)


## CSS

+ What's the box model?
+ How do you set margin, padding, width and height?

[CSS 框模型概述](http://www.w3school.com.cn/css/css_boxmodel.asp)

[盒子模型](http://zh.learnlayout.com/box-model.html)

+ What's the difference between `content-box` and `border-box`?

[box-sizing属性](http://sunyuhui.com/2015/03/30/box-sizing/)

## HTML

+ What's the relationship between DOM and HTML?
+ Do you know how to use the DOM in JavaScript?

Can you understand the following code?

```
<div id="container"></div>

<script>
var container = document.getElementById("container");

var n = 0;
setInterval(function() {
  n++;
  var newDiv = document.createElement("div");
  newDiv.innerHTML = n;
  container.appendChild(newDiv);
},1000);
</script>
```

See Demo: http://codepen.io/hayeah/pen/VvLyrp


+ [DOM概述](https://developer.mozilla.org/zh-CN/docs/Web/API/Document_Object_Model/Introduction)
+ [操作DOM](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/001434499851683f7f8d6b7717343248a75d5e7f930def4000)

# Git

If you are not familiar with Git, see: [Git Knowledge Checklistt](git/index.md)

# iOS

Your goal is to build a simple app in Xcode. First learn the Swift basics:

+ [Swift 入门教程 - Part 1](http://www.raywenderlich.com/74438/swift-tutorial-a-quick-start)

Then learn how to build a simple UI:

+ [Swift 入门教程 - Part 2](http://www.raywenderlich.com/74904/swift-tutorial-part-2-simple-ios-app)

The important things you need to know are:

1. UITextField, UILabel, UISlider are subclasses of UIView. They are the native UI elements in iOS. They are connected to the controller as `@IBOutlet`.
2. When the user touches something, an `@IBAction` function in the ViewController is called.
3. The ViewController set the properties of UIView objects to update the UI.
