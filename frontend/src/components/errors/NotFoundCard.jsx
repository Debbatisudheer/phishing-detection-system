import ErrorCard from "./ErrorCard";

function NotFoundCard() {

  return (

    <ErrorCard
      icon="❌"
      title="404"
      message="The page you requested does not exist."
      buttonText="Dashboard"
      onClick={() => {

        window.location.href="/";

      }}
    />

  );

}

export default NotFoundCard;