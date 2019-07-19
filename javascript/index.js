let socket = new WebSocket("ws://localhost:8080/ws")
let coins = 0;

console.log("Attempting Websocket Connection")

socket.onopen = () => {
    console.log("Successfuly Connected!")
    socket.send("Hi From The Client!")
}

socket.onclose = (event) => {
    console.log("Socket Close Connection: ", event)
}

socket.onmessage = (msg) => {
    console.log(msg);
}

socket.onerror = (error) => {
    console.log("Socket Error: ", error)
}

socket.addEventListener("message", function(e) {
    coins = msg;
    document.getElementById("bal").innerHTML = JSON.stringify(coins);
})