
var init = () => {
    alert('init body');    
var btn = document.getElementById("find");
var fxbt = () => {
        alert('whoaa you click a button');
    };
console.log(btn);    
btn.onclick=fxbt;
}