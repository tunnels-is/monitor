import ky from 'ky';
import { useEffect, useState } from 'react';
import { API_URL, DatacenterType } from './config/config';
import DatacenterComponent from './components/Datacenter';

const WS_URL = 'ws://localhost:1323/v1/json/dynamic';

function App() {
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
    });

    socket.addEventListener('error', (error) => {
      console.error('WebSocket error: ', error);
    });

    return () => {
      socket.close();
    };
  }, []);

  const [dataCenters, setDataCenters] = useState<Array<DatacenterType>>();

  useEffect(() => {
    const fetchData = async () => {
      const json = await ky.get(`${API_URL}/datacenters`).json();
      setDataCenters(json as Array<DatacenterType>);
    };
    console.log(dataCenters);
    fetchData();
  }, []);

  return (
    <div className="m-3">
      <h1 className="text-3xl my-3">Name?? Dashboard Stats Logo??</h1>
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
        {dataCenters?.map((item) => <DatacenterComponent dataCenter={item} />)}
      </div>
    </div>
  );
}

export default App;
