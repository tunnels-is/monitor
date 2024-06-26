import { RowType } from '@/config/config';

const Row = ({ Row }: { Row: RowType }) => {
  return (
    <div className="my-5 w-full">
      Row {Row.Tag}/{Row.ID}
      <div className="w-full h-10 border-2 border-gray-600 "></div>
    </div>
  );
};

export default Row;
