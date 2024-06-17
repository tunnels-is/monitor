import React, { useEffect, useState } from 'react';

const WS_URL = 'ws://localhost:1323/v1/json/dynamic';

function App() {
	const [message, setMessage] = useState('');
	const [socket, setSocket] = useState<WebSocket | null>(null);

	useEffect(() => {
		const socket = new WebSocket(WS_URL);
		setSocket(socket);

		socket.addEventListener('open', () => {
			console.log('Connected to server');
			socket.send('{"disk":0,"cpu":10,"memory":50,"sub":{"key1":1377}}');
		});

		socket.addEventListener('message', (event) => {
			console.log('Message from server ', event.data);
			setMessage(event.data);
		});

		socket.addEventListener('error', (error) => {
			console.error('WebSocket error: ', error);
		});

		return () => {
			socket.close();
		};
	}, []);

	const sendMessage = () => {
		if (socket && socket.readyState === WebSocket.OPEN) {
			socket.send('Hello World!');
		}
	};

	return (
		<div>
			<h1>WebSocket Example</h1>
			<p>Message from server: {message}</p>
			<button onClick={sendMessage}>Send Message</button>
		</div>
	);
}

export default App;
