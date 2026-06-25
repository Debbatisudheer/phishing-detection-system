import ErrorCard from "./ErrorCard";

function BackendOfflineCard() {

  return (

    <ErrorCard
      icon="⚠"
      title="Backend Unavailable"
      message="Unable to connect to the backend server."
      buttonText="Retry"
      onClick={() => window.location.reload()}
    />

  );

}

export default BackendOfflineCard;