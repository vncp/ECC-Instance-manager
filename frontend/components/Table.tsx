import React from "react";
import Table from "react-bootstrap/Table";
import Dropdown from "react-bootstrap/Dropdown";

export type Request = {
  name: string;
  netid: string;
  email: string;
  course: string;
  status: string;
};

const DataRow = ({ request, clickHandler }) => {
  return (
    <tr>
      <td>{request.name || "N/A"}</td>
      <td>{request.netid || "N/A"}</td>
      <td>{request.email || "N/A"}</td>
      <td>{request.course || "N/A"}</td>
      <td>{request.status || "N/A"}</td>
      <td>{request.date || "N/A"}</td>
      <td>
        <Dropdown>
          <Dropdown.Toggle id="dropdown-basic-button">Actions</Dropdown.Toggle>
          <Dropdown.Menu>
            <Dropdown.ItemText>
              <b>{request.netid}</b>
            </Dropdown.ItemText>
            <Dropdown.Item
              onClick={clickHandler}
              data-name={request.netid}
              data-task="createInstance"
            >
              Create Instance
            </Dropdown.Item>
            <Dropdown.Item
              onClick={clickHandler}
              data-name={request.netid}
              data-task="reloadInstance"
            >
              Restart Instance
            </Dropdown.Item>
            <Dropdown.Item
              onClick={clickHandler}
              data-name={request.netid}
              data-task="denyRequest"
            >
              Deny Request
            </Dropdown.Item>
          </Dropdown.Menu>
        </Dropdown>
      </td>
    </tr>
  );
};

const RequestTable = ({ requests, clickHandler }) => {
  let rows: JSX.Element[] = [];
  if (typeof requests == "object") {
    requests.forEach((r) => {
      rows.push(
        <DataRow request={r} key={r.netid} clickHandler={clickHandler} />
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
            <th>Email</th>
            <th>Course</th>
            <th>Status</th>
            <th>Date</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>{rows}</tbody>
      </Table>
    </div>
  );
};

export default RequestTable;
