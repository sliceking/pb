package actions

import (

  "fmt"
  "net/http"
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/pop/v5"
  "github.com/gobuffalo/x/responder"
  "pb/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Journal)
// DB Table: Plural (journals)
// Resource: Plural (Journals)
// Path: Plural (/journals)
// View Template Folder: Plural (/templates/journals/)

// JournalsResource is the resource for the Journal model
type JournalsResource struct{
  buffalo.Resource
}

// List gets all Journals. This function is mapped to the path
// GET /journals
func (v JournalsResource) List(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  journals := &models.Journals{}

  // Paginate results. Params "page" and "per_page" control pagination.
  // Default values are "page=1" and "per_page=20".
  q := tx.PaginateFromParams(c.Params())

  // Retrieve all Journals from the DB
  if err := q.All(journals); err != nil {
    return err
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // Add the paginator to the context so it can be used in the template.
    c.Set("pagination", q.Paginator)

    c.Set("journals", journals)
    return c.Render(http.StatusOK, r.HTML("/journals/index.plush.html"))
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(200, r.JSON(journals))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(200, r.XML(journals))
  }).Respond(c)
}

// Show gets the data for one Journal. This function is mapped to
// the path GET /journals/{journal_id}
func (v JournalsResource) Show(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Journal
  journal := &models.Journal{}

  // To find the Journal the parameter journal_id is used.
  if err := tx.Find(journal, c.Param("journal_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    c.Set("journal", journal)

    return c.Render(http.StatusOK, r.HTML("/journals/show.plush.html"))
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(200, r.JSON(journal))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(200, r.XML(journal))
  }).Respond(c)
}

// New renders the form for creating a new Journal.
// This function is mapped to the path GET /journals/new
func (v JournalsResource) New(c buffalo.Context) error {
  c.Set("journal", &models.Journal{})

  return c.Render(http.StatusOK, r.HTML("/journals/new.plush.html"))
}
// Create adds a Journal to the DB. This function is mapped to the
// path POST /journals
func (v JournalsResource) Create(c buffalo.Context) error {
  // Allocate an empty Journal
  journal := &models.Journal{}

  // Bind journal to the html form elements
  if err := c.Bind(journal); err != nil {
    return err
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(journal)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    return responder.Wants("html", func (c buffalo.Context) error {
      // Make the errors available inside the html template
      c.Set("errors", verrs)

      // Render again the new.html template that the user can
      // correct the input.
      c.Set("journal", journal)

      return c.Render(http.StatusUnprocessableEntity, r.HTML("/journals/new.plush.html"))
    }).Wants("json", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
    }).Wants("xml", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
    }).Respond(c)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a success message
    c.Flash().Add("success", T.Translate(c, "journal.created.success"))

    // and redirect to the show page
    return c.Redirect(http.StatusSeeOther, "/journals/%v", journal.ID)
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusCreated, r.JSON(journal))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusCreated, r.XML(journal))
  }).Respond(c)
}

// Edit renders a edit form for a Journal. This function is
// mapped to the path GET /journals/{journal_id}/edit
func (v JournalsResource) Edit(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Journal
  journal := &models.Journal{}

  if err := tx.Find(journal, c.Param("journal_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  c.Set("journal", journal)
  return c.Render(http.StatusOK, r.HTML("/journals/edit.plush.html"))
}
// Update changes a Journal in the DB. This function is mapped to
// the path PUT /journals/{journal_id}
func (v JournalsResource) Update(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Journal
  journal := &models.Journal{}

  if err := tx.Find(journal, c.Param("journal_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  // Bind Journal to the html form elements
  if err := c.Bind(journal); err != nil {
    return err
  }

  verrs, err := tx.ValidateAndUpdate(journal)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    return responder.Wants("html", func (c buffalo.Context) error {
      // Make the errors available inside the html template
      c.Set("errors", verrs)

      // Render again the edit.html template that the user can
      // correct the input.
      c.Set("journal", journal)

      return c.Render(http.StatusUnprocessableEntity, r.HTML("/journals/edit.plush.html"))
    }).Wants("json", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
    }).Wants("xml", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
    }).Respond(c)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a success message
    c.Flash().Add("success", T.Translate(c, "journal.updated.success"))

    // and redirect to the show page
    return c.Redirect(http.StatusSeeOther, "/journals/%v", journal.ID)
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.JSON(journal))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.XML(journal))
  }).Respond(c)
}

// Destroy deletes a Journal from the DB. This function is mapped
// to the path DELETE /journals/{journal_id}
func (v JournalsResource) Destroy(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Journal
  journal := &models.Journal{}

  // To find the Journal the parameter journal_id is used.
  if err := tx.Find(journal, c.Param("journal_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  if err := tx.Destroy(journal); err != nil {
    return err
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a flash message
    c.Flash().Add("success", T.Translate(c, "journal.destroyed.success"))

    // Redirect to the index page
    return c.Redirect(http.StatusSeeOther, "/journals")
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.JSON(journal))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.XML(journal))
  }).Respond(c)
}
