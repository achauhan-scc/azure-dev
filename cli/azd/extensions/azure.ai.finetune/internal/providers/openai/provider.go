// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package openai

import (
	"context"
<<<<<<< HEAD

	"azure.ai.finetune/internal/providers"
	"azure.ai.finetune/pkg/models"
)

// Ensure OpenAIProvider implements FineTuningProvider and ModelDeploymentProvider interfaces
var (
	_ providers.FineTuningProvider      = (*OpenAIProvider)(nil)
	_ providers.ModelDeploymentProvider = (*OpenAIProvider)(nil)
=======
	"fmt"
	"os"
	"time"

	"azure.ai.finetune/pkg/models"
	"github.com/azure/azure-dev/cli/azd/pkg/ux"
	"github.com/fatih/color"
	"github.com/openai/openai-go/v3"
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
)

// OpenAIProvider implements the provider interface for OpenAI APIs
type OpenAIProvider struct {
<<<<<<< HEAD
	// TODO: Add OpenAI SDK client
	// client *openai.Client
	apiKey   string
	endpoint string
}

// NewOpenAIProvider creates a new OpenAI provider instance
func NewOpenAIProvider(apiKey, endpoint string) *OpenAIProvider {
	return &OpenAIProvider{
		apiKey:   apiKey,
		endpoint: endpoint,
=======
	client *openai.Client
}

// NewOpenAIProvider creates a new OpenAI provider instance
func NewOpenAIProvider(client *openai.Client) *OpenAIProvider {
	return &OpenAIProvider{
		client: client,
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
	}
}

// CreateFineTuningJob creates a new fine-tuning job via OpenAI API
func (p *OpenAIProvider) CreateFineTuningJob(ctx context.Context, req *models.CreateFineTuningRequest) (*models.FineTuningJob, error) {
<<<<<<< HEAD
	// TODO: Implement
	// 1. Convert domain model to OpenAI SDK format
	// 2. Call OpenAI SDK CreateFineTuningJob
	// 3. Convert OpenAI response to domain model
	return nil, nil
=======

	params, err := convertInternalJobParamToOpenAiJobParams(req)
	if err != nil {
		return nil, fmt.Errorf("failed to convert internal model to openai: %w", err)
	}

	job, err := p.client.FineTuning.Jobs.New(ctx, *params)
	if err != nil {
		return nil, fmt.Errorf("failed to create fine-tuning job: %w", err)
	}

	return convertOpenAIJobToModel(*job), nil
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
}

// GetFineTuningStatus retrieves the status of a fine-tuning job
func (p *OpenAIProvider) GetFineTuningStatus(ctx context.Context, jobID string) (*models.FineTuningJob, error) {
	// TODO: Implement
	return nil, nil
}

// ListFineTuningJobs lists all fine-tuning jobs
func (p *OpenAIProvider) ListFineTuningJobs(ctx context.Context, limit int, after string) ([]*models.FineTuningJob, error) {
<<<<<<< HEAD
	// TODO: Implement
	return nil, nil
=======
	jobList, err := p.client.FineTuning.Jobs.List(ctx, openai.FineTuningJobListParams{
		Limit: openai.Int(int64(limit)), // optional pagination control
		After: openai.String(after),
	})

	if err != nil {
		return nil, err
	}

	var jobs []*models.FineTuningJob

	for _, job := range jobList.Data {
		finetuningJob := convertOpenAIJobToModel(job)
		jobs = append(jobs, finetuningJob)
	}

	return jobs, nil
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
}

// GetFineTuningJobDetails retrieves detailed information about a job
func (p *OpenAIProvider) GetFineTuningJobDetails(ctx context.Context, jobID string) (*models.FineTuningJobDetail, error) {
<<<<<<< HEAD
	// TODO: Implement
	return nil, nil
}

// GetJobEvents retrieves events for a fine-tuning job
func (p *OpenAIProvider) GetJobEvents(ctx context.Context, jobID string, limit int, after string) ([]*models.JobEvent, error) {
	// TODO: Implement
	return nil, nil
}

// GetJobCheckpoints retrieves checkpoints for a fine-tuning job
func (p *OpenAIProvider) GetJobCheckpoints(ctx context.Context, jobID string, limit int, after string) ([]*models.JobCheckpoint, error) {
	// TODO: Implement
	return nil, nil
=======
	job, err := p.client.FineTuning.Jobs.Get(ctx, jobID)
	if err != nil {
		return nil, err
	}
	finetuningJobDetail := convertOpenAIJobToDetailModel(job)

	return finetuningJobDetail, nil
}

// GetJobEvents retrieves events for a fine-tuning job
func (p *OpenAIProvider) GetJobEvents(ctx context.Context, jobID string) (*models.JobEventsList, error) {
	eventsPage, err := p.client.FineTuning.Jobs.ListEvents(
		ctx,
		jobID,
		openai.FineTuningJobListEventsParams{},
	)
	if err != nil {
		return nil, err
	}

	events := convertOpenAIJobEventsToModel(eventsPage)

	return events, nil
}

// GetJobCheckpoints retrieves checkpoints for a fine-tuning job
func (p *OpenAIProvider) GetJobCheckpoints(ctx context.Context, jobID string) (*models.JobCheckpointsList, error) {
	checkpointsPage, err := p.client.FineTuning.Jobs.Checkpoints.List(
		ctx,
		jobID,
		openai.FineTuningJobCheckpointListParams{},
	)
	if err != nil {
		return nil, err
	}
	checkpoints := convertOpenAIJobCheckpointsToModel(checkpointsPage)

	return checkpoints, nil
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
}

// PauseJob pauses a fine-tuning job
func (p *OpenAIProvider) PauseJob(ctx context.Context, jobID string) (*models.FineTuningJob, error) {
<<<<<<< HEAD
	// TODO: Implement
	return nil, nil
=======
	job, err := p.client.FineTuning.Jobs.Pause(ctx, jobID)
	if err != nil {
		return nil, err
	}

	finetuningJob := convertOpenAIJobToModel(*job)

	return finetuningJob, nil
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
}

// ResumeJob resumes a paused fine-tuning job
func (p *OpenAIProvider) ResumeJob(ctx context.Context, jobID string) (*models.FineTuningJob, error) {
<<<<<<< HEAD
	// TODO: Implement
	return nil, nil
=======
	job, err := p.client.FineTuning.Jobs.Resume(ctx, jobID)
	if err != nil {
		return nil, err
	}

	finetuningJob := convertOpenAIJobToModel(*job)

	return finetuningJob, nil
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
}

// CancelJob cancels a fine-tuning job
func (p *OpenAIProvider) CancelJob(ctx context.Context, jobID string) (*models.FineTuningJob, error) {
<<<<<<< HEAD
	// TODO: Implement
	return nil, nil
=======
	job, err := p.client.FineTuning.Jobs.Cancel(ctx, jobID)
	if err != nil {
		return nil, err
	}

	finetuningJob := convertOpenAIJobToModel(*job)

	return finetuningJob, nil
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
}

// UploadFile uploads a file for fine-tuning
func (p *OpenAIProvider) UploadFile(ctx context.Context, filePath string) (string, error) {
<<<<<<< HEAD
	// TODO: Implement
	return "", nil
=======
	if filePath == "" {
		return "", fmt.Errorf("file path cannot be empty")
	}

	// Show spinner while creating job
	spinner := ux.NewSpinner(&ux.SpinnerOptions{
		Text: "uploading the file for fine-tuning",
	})
	if err := spinner.Start(ctx); err != nil {
		fmt.Printf("failed to start spinner: %v\n", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		_ = spinner.Stop(ctx)
		return "", fmt.Errorf("\nfailed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	uploadedFile, err := p.client.Files.New(ctx, openai.FileNewParams{
		File:    file,
		Purpose: openai.FilePurposeFineTune,
	})

	if err != nil {
		_ = spinner.Stop(ctx)
		return "", fmt.Errorf("\nfailed to upload file: %w", err)
	}

	if uploadedFile == nil || uploadedFile.ID == "" {
		_ = spinner.Stop(ctx)
		return "", fmt.Errorf("\nuploaded file is empty")
	}

	// Poll for file processing status
	for {
		f, err := p.client.Files.Get(ctx, uploadedFile.ID)
		if err != nil {
			_ = spinner.Stop(ctx)
			return "", fmt.Errorf("\nfailed to check file status: %w", err)
		}
		if f.Status == openai.FileObjectStatusProcessed {
			_ = spinner.Stop(ctx)
			break
		}
		if f.Status == openai.FileObjectStatusError {
			_ = spinner.Stop(ctx)
			return "", fmt.Errorf("\nfile processing failed with status: %s", f.Status)
		}
		color.Yellow(".")
		time.Sleep(2 * time.Second)
	}

	return uploadedFile.ID, nil
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
}

// GetUploadedFile retrieves information about an uploaded file
func (p *OpenAIProvider) GetUploadedFile(ctx context.Context, fileID string) (interface{}, error) {
	// TODO: Implement
	return nil, nil
}

// DeployModel deploys a fine-tuned or base model
func (p *OpenAIProvider) DeployModel(ctx context.Context, req *models.DeploymentRequest) (*models.Deployment, error) {
	// TODO: Implement
	return nil, nil
}

// GetDeploymentStatus retrieves the status of a deployment
func (p *OpenAIProvider) GetDeploymentStatus(ctx context.Context, deploymentID string) (*models.Deployment, error) {
	// TODO: Implement
	return nil, nil
}

// ListDeployments lists all deployments
func (p *OpenAIProvider) ListDeployments(ctx context.Context, limit int, after string) ([]*models.Deployment, error) {
	// TODO: Implement
	return nil, nil
}

// UpdateDeployment updates deployment configuration
func (p *OpenAIProvider) UpdateDeployment(ctx context.Context, deploymentID string, capacity int32) (*models.Deployment, error) {
	// TODO: Implement
	return nil, nil
}

// DeleteDeployment deletes a deployment
func (p *OpenAIProvider) DeleteDeployment(ctx context.Context, deploymentID string) error {
	// TODO: Implement
	return nil
}
