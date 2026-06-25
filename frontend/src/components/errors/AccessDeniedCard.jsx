import ErrorCard from "./ErrorCard";

function AccessDeniedCard() {

  return (

    <ErrorCard
      icon="🚫"
      title="Access Denied"
      message="You don't have permission to access this page."
      buttonText="Home"
      onClick={() => window.location.href="/"}
    />

  );

}

export default AccessDeniedCard;