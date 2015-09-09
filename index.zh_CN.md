

# Sike.io ReactNative 训练营基础

React/ReactNative 训练营接触面比较广。你不需要对 web 开发有很深的了解，但是你确需要对杂七杂八的事情都知道一些。在训练营之前，我们准备了一个清单来帮助你学习 web 开发基础。

这个清单枚举了一系列的技术问题。如果你看到不理解的问题或者代码，请阅读我们链接到的学习资源。如果这些清单里的技术问题对你来说都很熟悉，你就已经具备了训练营需要的知识。

+ 如果你刚接触 web 开发，我们建议你应该尽可能地多花时间学习 JavaScript。
+ 如果你是一位前端开发者，你应该花些时间学习 iOS 基础。



# JavaScript

你能理解这段代码吗？



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



如果你对 JavaScript 语法不熟悉，阅读：

+ [快速入门](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/00143449917624134f5c4695b524e81a581ab5a222b05ec000)
+ [函数定义和调用](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/00143449926746982f181557d9b423f819e89709feabdb4000)

如果你对闭包和匿名函数不熟悉，阅读：

+ [闭包](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/00143449934543461c9d5dfeeb848f5b72bd012e1113d15000)
+ [变量作用域](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/0014344993159773a464f34e1724700a6d5dd9e235ceb7c000)



### JavaScript 中的对象

你能理解这段代码吗？



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



+ [方法](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/0014345005399057070809cfaa347dfb7207900cfd116fb000)
+ [创建对象](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/0014344997235247b53be560ab041a7b10360a567422a78000)
+ [原型继承](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/0014344997013405abfb7f0e1904a04ba6898a384b1e925000)



## CSS

+ 什么是盒模型？
+ 你如何设置 margin，padding，width 和 height？

[CSS 框模型概述](http://www.w3school.com.cn/css/css_boxmodel.asp)

[盒模型](http://zh.learnlayout.com/box-model.html)

+ `content-box` 和 `border-box` 之间的区别是什么？

[box-sizing属性](http://sunyuhui.com/2015/03/30/box-sizing/)



## HTML

+ DOM 和 HTML 的关系是什么？
+ 你知道如何在 JavaScript 中使用 DOM 吗？

你能理解下面的代码吗？



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



查看演示： http://codepen.io/hayeah/pen/VvLyrp


+ [DOM概述](https://developer.mozilla.org/zh-CN/docs/Web/API/Document_Object_Model/Introduction)
+ [操作DOM](http://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000/001434499851683f7f8d6b7717343248a75d5e7f930def4000)



# Git

如果你对 Git 不熟悉，阅读： [Git Knowledge Checklist](git/index.zh_CN.md)



# iOS

你的目标是在 Xcode 中开发一个简单的 app。首先, 学习 Swift 基础：

+ [Swift 入门教程 - Part 1](http://www.raywenderlich.com/74438/swift-tutorial-a-quick-start)

然后学习如何构建简单的 UI：

+ [Swift 入门教程 - Part 2](http://www.raywenderlich.com/74904/swift-tutorial-part-2-simple-ios-app)

你需要知道的重要的知识点：

1. UITextField， UILabel， UISlider 是 UIView 的子类。他们是 iOS 中的原生 UI 组件。它们通过 `@IBOutlet` 连接到 ViewController (控制器)。
2. 当用户触摸了控件，比如一个按钮，ViewController 中对应的 `@IBAction` 函数就会被调用。
3. ViewController 设置 UIView 对象的属性来更新 UI。

Xcode 百度盘镜像：https://github.com/iBcker/adcdownload
