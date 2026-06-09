import { useEffect, useState } from "react";
import api from "../services/api";

function CampaignTimeline() {

  const [timeline, setTimeline] =
    useState([]);

  useEffect(() => {

    loadTimeline();

  }, []);

  const loadTimeline =
    async () => {

      const token =
        localStorage.getItem(
          "token",
        );

      const response =
        await api.get(
          "/api/campaign-timeline",
          {
            headers: {
              Authorization:
                `Bearer ${token}`,
            },
          },
        );

      setTimeline(
        response.data,
      );
    };

  return (

    <div>

      <h1>
        Campaign Timeline
      </h1>

      <table
        border="1"
        width="100%"
      >

        <thead>

          <tr>

            <th>IOC</th>

            <th>First Seen</th>

            <th>Last Seen</th>

            <th>Occurrences</th>

          </tr>

        </thead>

        <tbody>

          {timeline.map(
            (item, index) => (

              <tr key={index}>

                <td>
                  {item.ioc}
                </td>

                <td>
                  {item.first_seen}
                </td>

                <td>
                  {item.last_seen}
                </td>

                <td>
                  {item.occurrences}
                </td>

              </tr>
            ),
          )}

        </tbody>

      </table>

    </div>
  );
}

export default CampaignTimeline;