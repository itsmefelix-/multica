package issueview

import (
	"strconv"

	"github.com/multica-ai/multica/server/internal/util"
	db "github.com/multica-ai/multica/server/pkg/db/generated"
)

// IssueResponse is the JSON response shape for an issue.
type IssueResponse struct {
	ID            string                  `json:"id"`
	WorkspaceID   string                  `json:"workspace_id"`
	Number        int32                   `json:"number"`
	Identifier    string                  `json:"identifier"`
	Title         string                  `json:"title"`
	Description   *string                 `json:"description"`
	Status        string                  `json:"status"`
	Priority      string                  `json:"priority"`
	AssigneeType  *string                 `json:"assignee_type"`
	AssigneeID    *string                 `json:"assignee_id"`
	CreatorType   string                  `json:"creator_type"`
	CreatorID     string                  `json:"creator_id"`
	ParentIssueID *string                 `json:"parent_issue_id"`
	ProjectID     *string                 `json:"project_id"`
	Position      float64                 `json:"position"`
	DueDate       *string                 `json:"due_date"`
	CreatedAt     string                  `json:"created_at"`
	UpdatedAt     string                  `json:"updated_at"`
	Reactions     []IssueReactionResponse `json:"reactions,omitempty"`
	Attachments   []AttachmentResponse    `json:"attachments,omitempty"`
}

type IssueReactionResponse struct {
	ID        string `json:"id"`
	IssueID   string `json:"issue_id"`
	ActorType string `json:"actor_type"`
	ActorID   string `json:"actor_id"`
	Emoji     string `json:"emoji"`
	CreatedAt string `json:"created_at"`
}

type AttachmentResponse struct {
	ID           string  `json:"id"`
	WorkspaceID  string  `json:"workspace_id"`
	IssueID      *string `json:"issue_id"`
	CommentID    *string `json:"comment_id"`
	UploaderType string  `json:"uploader_type"`
	UploaderID   string  `json:"uploader_id"`
	Filename     string  `json:"filename"`
	URL          string  `json:"url"`
	DownloadURL  string  `json:"download_url"`
	ContentType  string  `json:"content_type"`
	SizeBytes    int64   `json:"size_bytes"`
	CreatedAt    string  `json:"created_at"`
}

func IssueToResponse(i db.Issue, issuePrefix string) IssueResponse {
	identifier := issuePrefix + "-" + strconv.Itoa(int(i.Number))
	return IssueResponse{
		ID:            util.UUIDToString(i.ID),
		WorkspaceID:   util.UUIDToString(i.WorkspaceID),
		Number:        i.Number,
		Identifier:    identifier,
		Title:         i.Title,
		Description:   util.TextToPtr(i.Description),
		Status:        i.Status,
		Priority:      i.Priority,
		AssigneeType:  util.TextToPtr(i.AssigneeType),
		AssigneeID:    util.UUIDToPtr(i.AssigneeID),
		CreatorType:   i.CreatorType,
		CreatorID:     util.UUIDToString(i.CreatorID),
		ParentIssueID: util.UUIDToPtr(i.ParentIssueID),
		ProjectID:     util.UUIDToPtr(i.ProjectID),
		Position:      i.Position,
		DueDate:       util.TimestampToPtr(i.DueDate),
		CreatedAt:     util.TimestampToString(i.CreatedAt),
		UpdatedAt:     util.TimestampToString(i.UpdatedAt),
	}
}

func IssueListRowToResponse(i db.ListIssuesRow, issuePrefix string) IssueResponse {
	identifier := issuePrefix + "-" + strconv.Itoa(int(i.Number))
	return IssueResponse{
		ID:            util.UUIDToString(i.ID),
		WorkspaceID:   util.UUIDToString(i.WorkspaceID),
		Number:        i.Number,
		Identifier:    identifier,
		Title:         i.Title,
		Status:        i.Status,
		Priority:      i.Priority,
		AssigneeType:  util.TextToPtr(i.AssigneeType),
		AssigneeID:    util.UUIDToPtr(i.AssigneeID),
		CreatorType:   i.CreatorType,
		CreatorID:     util.UUIDToString(i.CreatorID),
		ParentIssueID: util.UUIDToPtr(i.ParentIssueID),
		ProjectID:     util.UUIDToPtr(i.ProjectID),
		Position:      i.Position,
		DueDate:       util.TimestampToPtr(i.DueDate),
		CreatedAt:     util.TimestampToString(i.CreatedAt),
		UpdatedAt:     util.TimestampToString(i.UpdatedAt),
	}
}

func OpenIssueRowToResponse(i db.ListOpenIssuesRow, issuePrefix string) IssueResponse {
	identifier := issuePrefix + "-" + strconv.Itoa(int(i.Number))
	return IssueResponse{
		ID:            util.UUIDToString(i.ID),
		WorkspaceID:   util.UUIDToString(i.WorkspaceID),
		Number:        i.Number,
		Identifier:    identifier,
		Title:         i.Title,
		Status:        i.Status,
		Priority:      i.Priority,
		AssigneeType:  util.TextToPtr(i.AssigneeType),
		AssigneeID:    util.UUIDToPtr(i.AssigneeID),
		CreatorType:   i.CreatorType,
		CreatorID:     util.UUIDToString(i.CreatorID),
		ParentIssueID: util.UUIDToPtr(i.ParentIssueID),
		ProjectID:     util.UUIDToPtr(i.ProjectID),
		Position:      i.Position,
		DueDate:       util.TimestampToPtr(i.DueDate),
		CreatedAt:     util.TimestampToString(i.CreatedAt),
		UpdatedAt:     util.TimestampToString(i.UpdatedAt),
	}
}
