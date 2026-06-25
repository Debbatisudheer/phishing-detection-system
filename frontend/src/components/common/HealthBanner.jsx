import {
  useEffect,
  useState,
} from "react";

import api
from "../../services/api";

function HealthBanner() {

  const [healthy,
    setHealthy] =
      useState(true);

  const checkHealth =
    async () => {

      try {

        await api.get(
          "/health",
        );

        setHealthy(true);

      } catch {

        setHealthy(false);
      }
    };

  useEffect(() => {

    checkHealth();

    const interval =
      setInterval(
        checkHealth,
        30000,
      );

    return () =>
      clearInterval(
        interval,
      );

  }, []);

  return (

    <div
      style={{
        padding: "10px",
        textAlign: "center",
      }}
    >

      {

        healthy

          ?

          "🟢 All Systems Operational"

          :

          "🔴 Backend Offline"

      }

    </div>

  );
}

export default HealthBanner;