import ErrorCard from "./ErrorCard";

function SandboxTimeoutCard() {

  return (

    <ErrorCard
      title="⏳ Sandbox Timeout"
      message="
Sandbox analysis exceeded
the maximum execution time.
"
    />

  );
}

export default SandboxTimeoutCard;