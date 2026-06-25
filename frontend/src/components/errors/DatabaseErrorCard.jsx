import ErrorCard from "./ErrorCard";

function DatabaseErrorCard() {

  return (

    <ErrorCard
      title="🗄 Database Error"
      message="
Unable to retrieve data
from database.
"
    />

  );
}

export default DatabaseErrorCard;