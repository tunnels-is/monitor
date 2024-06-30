import { API_URL, DatacenterType, RowType } from '@/config/config';
import ky from 'ky';
import { useEffect, useState } from 'react';
import Row from './Row';

const Datacenter = ({ dataCenter }: { dataCenter: DatacenterType }) => {
  const [rows, setRows] = useState<RowType[]>();

  useEffect(() => {
    const getRows = async () => {
      const json = await ky
        .get(`${API_URL}/rows/datacenter/${dataCenter.ID}`)
        .json();
      setRows(json as RowType[]);
      console.log(json);
    };
    getRows();
  }, []);

  return (
    <div className="m-3">
      Datacenter {dataCenter.Tag}/{dataCenter.ID}
      <div className="w-fit p-3 min-h-56 min-w-72 w-full border-2 border-gray-600">
        {rows?.map((row) => <Row Row={row} />)}
      </div>
    </div>
  );
};

export default Datacenter;
