import React, { useState, useEffect } from "react";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogTitle from "@mui/material/DialogTitle";

export default function FormDialog(props) {
  const [formValues, setFormValues] = useState(props.valueData);



  const handleCloseDialog = () => {
    props.setOpen(false);
  };

  const handleFormSubmit = (event) => {
    event.preventDefault();

    props.handleSubmit(formValues,props.type);
    handleCloseDialog();
  };

  const handleChange = (event) => {
    const { name, value } = event.target;
    setFormValues((prevState) => ({
      ...prevState,
      [name]: value,
    }));
  };

  return (
    <Dialog
      open={props.open}
      onClose={handleCloseDialog}
      PaperProps={{
        component: "form",
        onSubmit: handleFormSubmit,
      }}
    >
      <DialogTitle>Update</DialogTitle>
      <DialogContent>
        <DialogContentText>{props.contentText}</DialogContentText>
        {Object.entries(formValues).map(([fieldName, fieldValue]) => {
          if (fieldName === "taskId" || fieldName === "isComplete"|| fieldName === "subtaskId"||(fieldName === "subtasks" && props.type === "task")) {
            // Skip rendering this field
            return null;
          }
          return (
            <TextField
              key={fieldName}
              required
              name={fieldName}
              label={fieldName}
              type="text"
              variant="standard"
              value={fieldValue}
              onChange={handleChange}
            />
          );
        })}
        
      </DialogContent>
      <DialogActions>
        <Button onClick={handleCloseDialog}>Cancel</Button>
        <Button type="submit">{props.submitButtonText}</Button>
      </DialogActions>
    </Dialog>
  );
}
