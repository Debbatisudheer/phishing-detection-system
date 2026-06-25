import ErrorCard from "./ErrorCard";

function EmptyStateCard({

  title,

  message,

}) {

  return (

    <ErrorCard
      icon="📂"
      title={title}
      message={message}
      buttonText="Refresh"
      onClick={() => window.location.reload()}
    />

  );

}

export default EmptyStateCard;