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
