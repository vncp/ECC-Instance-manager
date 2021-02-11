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

const DataRow = ({ request }) => {
  return (
    <tr>
      <td>{request.name}</td>
      <td>{request.netid}</td>
      <td>{request.email}</td>
      <td>{request.course}</td>
      <td>{request.status}</td>
      <td>{request.date}</td>
      <td>
        <Dropdown>
          <Dropdown.Toggle id="dropdown-basic-button">Actions</Dropdown.Toggle>
          <Dropdown.Menu>
            <Dropdown.ItemText>
              <b>{request.netid}</b>
            </Dropdown.ItemText>
            <Dropdown.Item data-name={request.netid} data-task="createInstance">
              Create Instance
            </Dropdown.Item>
            <Dropdown.Item data-name={request.netid} data-task="reloadInstance">
              Restart Instance
            </Dropdown.Item>
            <Dropdown.Item data-name={request.netid} data-task="denyRequest">
              Deny Request
            </Dropdown.Item>
          </Dropdown.Menu>
        </Dropdown>
      </td>
    </tr>
  );
};

const RequestTable = ({ requests }) => {
  let rows: JSX.Element[] = [];
  requests.forEach((r) => {
    rows.push(<DataRow request={r} key={r.netid} />);
  });
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
