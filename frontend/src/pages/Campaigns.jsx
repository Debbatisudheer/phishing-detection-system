import { useEffect, useState } from "react";
import api from "../services/api";

function Campaigns() {

  const [campaigns, setCampaigns] =
    useState([]);

  const [sources, setSources] =
    useState([]);

  useEffect(() => {

    loadCampaigns();

  }, []);

  const loadCampaigns =
    async () => {

      const token =
        localStorage.getItem(
          "token",
        );

      const response =
        await api.get(
          "/api/campaigns",
          {
            headers: {
              Authorization:
                `Bearer ${token}`,
            },
          },
        );

      setCampaigns(
        response.data,
      );
    };

  const loadSources =
    async (ioc) => {

      const token =
        localStorage.getItem(
          "token",
        );

      const response =
        await api.get(
          `/api/ioc-sources?ioc=${encodeURIComponent(ioc)}`,
          {
            headers: {
              Authorization:
                `Bearer ${token}`,
            },
          },
        );

      setSources(
        response.data,
      );
    };

  return (

    <div>

      <h1>
        Campaign Dashboard
      </h1>

      <table
        border="1"
        width="100%"
      >

        <thead>

          <tr>

            <th>IOC</th>

            <th>Count</th>

            <th>Severity</th>

            <th>Action</th>

          </tr>

        </thead>

        <tbody>

          {campaigns.map(
            (campaign, index) => (

              <tr key={index}>

                <td>
                  {campaign.ioc}
                </td>

                <td>
                  {campaign.count}
                </td>

                <td>
                  {campaign.severity}
                </td>

                <td>

                  <button
                    onClick={() =>
                      loadSources(
                        campaign.ioc,
                      )
                    }
                  >
                    View Sources
                  </button>

                </td>

              </tr>
            ),
          )}

        </tbody>

      </table>

      <h2>
        IOC Sources
      </h2>

      <table
        border="1"
        width="100%"
      >

        <thead>

          <tr>

            <th>Source</th>

            <th>File</th>

            <th>Time</th>

          </tr>

        </thead>

        <tbody>

          {sources.map(
            (source, index) => (

              <tr key={index}>

                <td>
                  {source.source_type}
                </td>

                <td>
                  {source.file_name}
                </td>

                <td>
                  {source.created_at}
                </td>

              </tr>
            ),
          )}

        </tbody>

      </table>

    </div>
  );
}

export default Campaigns;