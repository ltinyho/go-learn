/**
 * 四则运算生成逆波兰表达式
 * @constructor
 */
function Stack() {
  const data = [];
  this.push = push;
  this.pop = pop;
  this.peek = peek;
  this.print = print;
  this.getData = getData;

  function push(element) {
    data.push(element);
  }

  function pop() {
    return data.pop();
  }

  function peek() {
    return data[data.length - 1];
  }

  function print() {
    console.log(data);
  }

  function getData() {
    return data;
  }
}

let infix = '9 + ( 3 - 1 ) * 3 + 10 / 2';
infix = infix.split(' '); // 将表达式分隔开,方便遍历
const s = new Stack();
const ops = '-+/*'; // 定义操作符
const precedence = { '*': 3, '/': 3, '+': 2, '-': 2 }; // 定义运算符的优先级
let token;
let postfix = []; // 结果

for (let i = 0; i < infix.length; i++) {
  token = infix[i];
  if (!isNaN(Number(token))) { // if token is operand
    postfix.push(token);
  } else if (ops.indexOf(token) !== -1) { // if token is an operator
    let o2 = s.peek();
    while (ops.indexOf(o2) !== -1 && precedence[token] <= precedence[o2]) {
      postfix.push(s.pop());
      o2 = s.peek(); // next round\
    }
    s.push(token); // push token onto the stack
  } else if (token === '(') { // if token is left parenthesis
    s.push(token); // then push it onto the stack
  } else if (token === ')') { // if token is right parenthesis
    // 遇到右括号,将操作数栈弹出,直到遇到左括号
    while (s.peek() !== '(') { // until token at top is (
      const top = s.pop();
      postfix.push(top);
    }
    s.pop(); // 移除 ( ,但不推到结果中
  }
}
// 剩余的操作符反转,推送到结果
postfix = postfix.concat(s.getData().reverse()).join(' ');
console.log(postfix);
