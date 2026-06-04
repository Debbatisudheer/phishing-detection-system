import { useEffect, useState } from "react";
import api from "../services/api";

function Dashboard() {

  const [stats, setStats] = useState({});

  useEffect(() => {

    const fetchDashboard = async () => {

      const token =
        localStorage.getItem(
          "token",
        );

      const response =
        await api.get(
          "/api/dashboard",
          {
            headers: {
              Authorization:
                `Bearer ${token}`,
            },
          },
        );

      setStats(
        response.data,
      );
    };

    fetchDashboard();

  }, []);

  return (
    <div>

      <h1>
        SOC Dashboard
      </h1>

      <h3>
        Total Analyzed:
        {stats.total_analyzed}
      </h3>

      <h3>
        Allow:
        {stats.allow}
      </h3>

      <h3>
        Suspicious:
        {stats.suspicious}
      </h3>

      <h3>
        Quarantine:
        {stats.quarantine}
      </h3>

      <h3>
        Critical:
        {stats.critical}
      </h3>

    </div>
  );
}

export default Dashboard;