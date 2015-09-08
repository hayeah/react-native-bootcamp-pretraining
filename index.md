# Sike.io ReactNative Bootcamp Fundamentals

To get ready for the React/ReactNative bootcamp, you don't need to know web development very well, but you do need to know a little bit about a lot of things. We prepared this checklist to help you learn web development fundamentals before we start the bootcamp.

This checklist is a list of questions. If you see a question or a piece of code that you don't understand, read the learning resources we've linked to. If everything in this checklist is familiar to you, then you are ready for the bootcamp.

+ If you are new to web development, we recommend that you should spend as much time as you can learning JavaScript.
+ If you are a frontend developer, you should spend some time learning iOS fundamentals.

<cn>

# Sike.io ReactNative 训练营基础

React/ReactNative 训练营接触面比较广。你不需要对 web 开发有很深的了解，但是你确需要对杂七杂八的事情都知道一些。在训练营之前，我们准备了一个清单来帮助你学习 web 开发基础。

这个清单枚举了一系列的技术问题。如果你看到不理解的问题或者代码，请阅读我们链接到的学习资源。如果这些清单里的技术问题对你来说都很熟悉，你就已经具备了训练营需要的知识。

+ 如果你刚接触 web 开发，我们建议你应该尽可能地多花时间学习 JavaScript。
+ 如果你是一位前端开发者，你应该花些时间学习 iOS 基础。

</cn>

# JavaScript

Can you understand this piece of code?

<cn>

# JavaScript

你能理解这段代码吗？

</cn>

```js
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


  // Do you know that `cancel` is a closure that captures the timerID?
  var cancel = function() {
    clearInterval(timerID);
  };

  // Do you know that closure/function is a value?
  return cancel;
}

// Do you know what a callback function is?
var cancel = startCounter(500,function(i) {
  console.log("current counter value:",i);
});

// stop the counter after 5 seconds.
setTimeout(cancel,5000);


// Output
// current counter value: 1
// current counter value: 2
// current counter value: 3
// current counter value: 4
// current counter value: 5
// current counter value: 6
// current counter value: 7
// current counter value: 8
// current counter value: 9

```

<cn>

```js
/**
* 这是一个函数，它在固定间隔调用回调函数。
*
* @param interval 在每次计数器增长之间的时间间隔。
* @param callback 在每个间隔对回调函数传递计数器的值。
* @return 一个可以取消计数器的函数。
*/
function startCounter(interval,callback) {
  var i = 0;

  // 你知道匿名函数是什么吗？
  var timerID = setInterval(function() {
    i++;
    callback(i);
  },interval);


  // 你知道 `cancel` 是一个闭包，它捕获了 timeID 吗？
  var cancel = function() {
    clearInterval(timerID);
  };

  // 你知道闭包/函数是一个值吗？
  return cancel;
}

// 你知道什么是回调函数吗？
var cancel = startCounter(500,function(i) {
  console.log("当前计数器的值：",i);
});

// 5 秒后停止计数器
setTimeout(cancel,5000);


// 输出
// 当前计数器的值：1
// 当前计数器的值：2
// 当前计数器的值：3
// 当前计数器的值：4
// 当前计数器的值：5
// 当前计数器的值：6
// 当前计数器的值：7
// 当前计数器的值：8
// 当前计数器的值：9

```

</cn>

If you are not familiar with JavaScript syntax, read:

+ [快速入门](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/00143449917624134f5c4695b524e81a581ab5a222b05ec000)
+ [函数定义和调用](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/00143449926746982f181557d9b423f819e89709feabdb4000)

If you are not familiar with closure and anonymous function, read:

+ [闭包](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/00143449934543461c9d5dfeeb848f5b72bd012e1113d15000)
+ [变量作用域](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/0014344993159773a464f34e1724700a6d5dd9e235ceb7c000)

<cn>

如果你对 JavaScript 语法不熟悉，阅读：

+ [快速入门](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/00143449917624134f5c4695b524e81a581ab5a222b05ec000)
+ [函数定义和调用](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/00143449926746982f181557d9b423f819e89709feabdb4000)

如果你对闭包和匿名函数不熟悉，阅读：

+ [闭包](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/00143449934543461c9d5dfeeb848f5b72bd012e1113d15000)
+ [变量作用域](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/0014344993159773a464f34e1724700a6d5dd9e235ceb7c000)

</cn>

### Objects in JavaScript

Can you understand this piece of code?

<cn>

### JavaScript 中的对象

你能理解这段代码吗？

</cn>

```js
// Person is a constructor
function Person(name) {
  this.name;
}

// Every instance of Person would have the `sayHello` method.
Person.prototype.sayHello = function() {
  console.log("hello, my name is " + this.name);
}

// Create an instance of Person using the Person constructor.
var person = new Person();

// Calls person's method.
person.sayHello();
```

<cn>

```js
// Person 是一个构造函数
function Person(name) {
  this.name;
}

// 每个 Person 实例会有 `sayHello` 方法。
Person.prototype.sayHello = function() {
  console.log("你好，我的名字是 " + this.name);
}

// 使用 Person 构造函数创建一个 Person 实例。
var person = new Person();

// 调用 person 的方法。
person.sayHello();
```

</cn>

+ [方法](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/0014345005399057070809cfaa347dfb7207900cfd116fb000)
+ [创建对象](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/0014344997235247b53be560ab041a7b10360a567422a78000)
+ [原型继承](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/0014344997013405abfb7f0e1904a04ba6898a384b1e925000)

<cn>

+ [方法](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/0014345005399057070809cfaa347dfb7207900cfd116fb000)
+ [创建对象](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/0014344997235247b53be560ab041a7b10360a567422a78000)
+ [原型继承](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/0014344997013405abfb7f0e1904a04ba6898a384b1e925000)

</cn>


## CSS

+ What's the box model?
+ How do you set margin, padding, width and height?

[CSS 框模型概述](http://www.w3school.com.cn/css/css_boxmodel.asp)

[盒子模型](http://zh.learnlayout.com/box-model.html)

+ What's the difference between `content-box` and `border-box`?

[box-sizing属性](http://sunyuhui.com/2015/03/30/box-sizing/)

<cn>

## CSS

+ 什么是盒模型？
+ 你如何设置 margin，padding，width 和 height？

[CSS 框模型概述](http://www.w3school.com.cn/css/css_boxmodel.asp)

[盒模型](http://zh.learnlayout.com/box-model.html)

+ `content-box` 和 `border-box` 之间的区别是什么？

[box-sizing属性](http://sunyuhui.com/2015/03/30/box-sizing/)

</cn>

## HTML

+ What's the relationship between DOM and HTML?
+ Do you know how to use the DOM in JavaScript?

Can you understand the following code?

<cn>

## HTML

+ DOM 和 HTML 的关系是什么？
+ 你知道如何在 JavaScript 中使用 DOM 吗？

你能理解下面的代码吗？

</cn>

```html
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

<cn>

```html
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

</cn>

See Demo: http://codepen.io/hayeah/pen/VvLyrp


+ [DOM概述](https://developer.mozilla.org/zh-CN/docs/Web/API/Document_Object_Model/Introduction)
+ [操作DOM](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/001434499851683f7f8d6b7717343248a75d5e7f930def4000)

<cn>

查看演示： http://codepen.io/hayeah/pen/VvLyrp


+ [DOM概述](https://developer.mozilla.org/zh-CN/docs/Web/API/Document_Object_Model/Introduction)
+ [操作DOM](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/001434499851683f7f8d6b7717343248a75d5e7f930def4000)

</cn>

# Git

If you are not familiar with Git, see: [Git Knowledge Checklist](git/index.md)

<cn>

# Git

如果你对 Git 不熟悉，阅读： [Git Knowledge Checklist](git/index.zh_CN.md)

</cn>

# iOS

Your goal is to build a simple app in Xcode. First, learn the Swift basics:

+ [Swift 入门教程 - Part 1](http://www.raywenderlich.com/74438/swift-tutorial-a-quick-start)

Then learn how to build a simple UI:

+ [Swift 入门教程 - Part 2](http://www.raywenderlich.com/74904/swift-tutorial-part-2-simple-ios-app)

The important ideas you need to know are:

1. UITextField, UILabel, UISlider are subclasses of UIView. They are the native UI components in iOS. They are connected to the controller as `@IBOutlet`.
2. When the user touches a component (a button, or a slider), the corresponding `@IBAction` function in the ViewController would be called.
3. The ViewController sets the properties of UIView objects to update the UI.

<cn>

# iOS

你的目标是在 Xcode 中开发一个简单的 app。首先, 学习 Swift 基础：

+ [Swift 入门教程 - Part 1](http://www.raywenderlich.com/74438/swift-tutorial-a-quick-start)

然后学习如何构建简单的 UI：

+ [Swift 入门教程 - Part 2](http://www.raywenderlich.com/74904/swift-tutorial-part-2-simple-ios-app)

你需要知道的重要的知识点：

1. UITextField， UILabel， UISlider 是 UIView 的子类。他们是 iOS 中的原生 UI 组件。它们通过 `@IBOutlet` 连接到 ViewController (控制器)。
2. 当用户触摸了控件，比如一个按钮，ViewController 中对应的 `@IBAction` 函数就会被调用。
3. ViewController 设置 UIView 对象的属性来更新 UI。

</cn>