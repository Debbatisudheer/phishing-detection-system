import { useEffect, useState }
from "react";

function LiveAlerts() {

  const [alerts, setAlerts] =
    useState([]);

  useEffect(() => {

    const socket =
      new WebSocket(
        "ws://localhost:8081/ws",
      );

    socket.onmessage =
      (event) => {

        setAlerts(
          (prev) => [
            event.data,
            ...prev,
          ],
        );
      };

    return () => {

      socket.close();
    };

  }, []);

  return (

    <div>

      <h1>
        Live Alerts
      </h1>

      {alerts.map(
        (alert, index) => (

          <div
            key={index}
  className="card alert-critical"
            style={{
              border:
                "1px solid red",
              padding:
                "10px",
              margin:
                "10px",
            }}
          >

            {alert}

          </div>
        ),
      )}

    </div>
  );
}

export default LiveAlerts;