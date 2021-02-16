import React from "react";
import Table from "react-bootstrap/Table";
import Dropdown from "react-bootstrap/Dropdown";

export type Instance = {
  name: string;
  netid: string;
  status: string;
};

const DataRow = ({ instance, clickHandler, staff }) => {
  return (
    <tr>
      <td>{instance.name || "N/A"}</td>
      <td>{instance.netid || "N/A"}</td>
      <td>{instance.status || "N/A"}</td>
      <td>
        <Dropdown>
          <Dropdown.Toggle id="dropdown-basic-button">Actions</Dropdown.Toggle>
          <Dropdown.Menu>
            <Dropdown.ItemText>
              <b>{instance.netid}</b>
            </Dropdown.ItemText>
            <Dropdown.Item
              onClick={clickHandler}
              data-name={instance.netid}
              data-task="reloadInstance"
            >
              Restart Instance
            </Dropdown.Item>
            {staff ? (
              <>
                <Dropdown.Item
                  onClick={clickHandler}
                  data-name={instance.netid}
                  data-task="createInstance"
                >
                  Recreate Instance
                </Dropdown.Item>
                <Dropdown.Item
                  onClick={clickHandler}
                  data-name={instance.netid}
                  data-task="deleteInstance"
                >
                  Delete Instance
                </Dropdown.Item>
              </>
            ) : (
              <></>
            )}
          </Dropdown.Menu>
        </Dropdown>
      </td>
    </tr>
  );
};

const ManagementTable = ({ instances, clickHandler, staff }) => {
  let rows: JSX.Element[] = [];
  if (typeof instances == "object") {
    instances.forEach((i) => {
      rows.push(
        <DataRow
          instance={i}
          key={i.netid}
          clickHandler={clickHandler}
          staff={staff}
        />
      );
    });
  }
  return (
    <div>
      <Table bordered striped>
        <thead>
          <tr>
            <th>Name</th>
            <th>Netid</th>
            <th>Status</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>{rows}</tbody>
      </Table>
    </div>
  );
};

export default ManagementTable;
