import { useEffect, useState } from "react";
import ReactFlow from "reactflow";
import "reactflow/dist/style.css";
import api from "../services/api";

function IOCNetworkGraph() {

  const [elements, setElements] =
    useState({
      nodes: [],
      edges: [],
    });

  useEffect(() => {

    loadGraph();

  }, []);

  const loadGraph =
    async () => {

      const token =
        localStorage.getItem(
          "token",
        );

      const response =
        await api.get(
          "/api/ioc-graph",
          {
            headers: {
              Authorization:
                `Bearer ${token}`,
            },
          },
        );

      const nodes = [];
      const edges = [];

      const addedNodes =
        new Set();

      response.data.forEach(
        (
          item,
          index,
        ) => {

          const iocId =
            `ioc-${item.ioc}`;

          const sourceId =
            `source-${item.source}-${index}`;

          if (
            !addedNodes.has(
              iocId,
            )
          ) {

            nodes.push({

              id: iocId,

              data: {
                label:
                  item.ioc,
              },

              position: {
                x: 400,
                y:
                  index *
                  120,
              },
            });

            addedNodes.add(
              iocId,
            );
          }

          nodes.push({

            id: sourceId,

            data: {
              label:
                `${item.source}
                 (${item.file})`,
            },

            position: {
              x: 50,
              y:
                index *
                120,
            },
          });

          edges.push({

            id:
              `edge-${index}`,

            source:
              sourceId,

            target:
              iocId,
          });
        },
      );

      setElements({
        nodes,
        edges,
      });
    };

  return (

    <div
      style={{
        height: "700px",
      }}
    >

      <h1>
        IOC Relationship Network
      </h1>

      <ReactFlow
        nodes={
          elements.nodes
        }
        edges={
          elements.edges
        }
        fitView
      />

    </div>
  );
}

export default IOCNetworkGraph;