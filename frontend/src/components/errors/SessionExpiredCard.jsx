import ErrorCard from "./ErrorCard";

function SessionExpiredCard() {

  return (

    <ErrorCard
      icon="🔐"
      title="Session Expired"
      message="Your session has expired. Please login again."
      buttonText="Login"
      onClick={() => {

        localStorage.removeItem("token");

        window.location.href="/login";

      }}
    />

  );

}

export default SessionExpiredCard;