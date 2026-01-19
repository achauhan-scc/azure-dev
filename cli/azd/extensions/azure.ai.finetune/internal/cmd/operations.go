// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package cmd

import (
	"fmt"
	"strings"

<<<<<<< HEAD
	"github.com/azure/azure-dev/cli/azd/pkg/azdext"
	"github.com/azure/azure-dev/cli/azd/pkg/ux"
	"github.com/fatih/color"

	"github.com/spf13/cobra"

	FTYaml "azure.ai.finetune/internal/fine_tuning_yaml"
	JobWrapper "azure.ai.finetune/internal/tools"
=======
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/azure/azure-dev/cli/azd/pkg/azdext"
	"github.com/azure/azure-dev/cli/azd/pkg/ux"

	"azure.ai.finetune/internal/services"
	"azure.ai.finetune/internal/utils"
	"azure.ai.finetune/pkg/models"
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
)

func newOperationCommand() *cobra.Command {
	cmd := &cobra.Command{
<<<<<<< HEAD
		Use:   "jobs",
=======
		Use: "jobs",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			return validateEnvironment(cmd.Context())
		},
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
		Short: "Manage fine-tuning jobs",
	}

	cmd.AddCommand(newOperationSubmitCommand())
	cmd.AddCommand(newOperationShowCommand())
	cmd.AddCommand(newOperationListCommand())
<<<<<<< HEAD
	cmd.AddCommand(newOperationActionCommand())
	cmd.AddCommand(newOperationDeployModelCommand())
=======
	cmd.AddCommand(newOperationPauseCommand())
	cmd.AddCommand(newOperationResumeCommand())
	cmd.AddCommand(newOperationCancelCommand())
	// cmd.AddCommand(newOperationDeployModelCommand())
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7

	return cmd
}

<<<<<<< HEAD
// getStatusSymbol returns a symbol representation for job status
func getStatusSymbol(status string) string {
	switch status {
	case "pending":
		return "⌛"
	case "queued":
		return "📚"
	case "running":
		return "🔄"
	case "succeeded":
		return "✅"
	case "failed":
		return "💥"
	case "cancelled":
		return "❌"
	default:
		return "❓"
	}
}

// formatFineTunedModel returns the model name or "NA" if blank
func formatFineTunedModel(model string) string {
	if model == "" {
		return "NA"
	}
	return model
}

func newOperationSubmitCommand() *cobra.Command {
	var filename string
	cmd := &cobra.Command{
		Use:   "submit",
		Short: "Submit fine tuning job",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := azdext.WithAccessToken(cmd.Context())

			// Validate filename is provided
			if filename == "" {
				return fmt.Errorf("config file is required, use -f or --file flag")
=======
func newOperationSubmitCommand() *cobra.Command {
	var filename string
	var model string
	var trainingFile string
	var validationFile string
	var suffix string
	var seed int64
	cmd := &cobra.Command{
		Use:   "submit",
		Short: "submit fine tuning job",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := azdext.WithAccessToken(cmd.Context())
			if filename == "" && (model == "" || trainingFile == "") {
				return fmt.Errorf("either config file or model and training-file parameters are required")
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
			}

			azdClient, err := azdext.NewAzdClient()
			if err != nil {
				return fmt.Errorf("failed to create azd client: %w", err)
			}
			defer azdClient.Close()

<<<<<<< HEAD
			// Parse and validate the YAML configuration file
			color.Green("Parsing configuration file...")
			config, err := FTYaml.ParseFineTuningConfig(filename)
			if err != nil {
				return err
			}

			// Upload training file

			trainingFileID, err := JobWrapper.UploadFileIfLocal(ctx, azdClient, config.TrainingFile)
			if err != nil {
				return fmt.Errorf("failed to upload training file: %w", err)
			}

			// Upload validation file if provided
			var validationFileID string
			if config.ValidationFile != "" {
				validationFileID, err = JobWrapper.UploadFileIfLocal(ctx, azdClient, config.ValidationFile)
				if err != nil {
					return fmt.Errorf("failed to upload validation file: %w", err)
				}
			}

			// Create fine-tuning job
			// Convert YAML configuration to OpenAI job parameters
			jobParams, err := ConvertYAMLToJobParams(config, trainingFileID, validationFileID)
			if err != nil {
				return fmt.Errorf("failed to convert configuration to job parameters: %w", err)
			}

			// Submit the fine-tuning job using CreateJob from JobWrapper
			job, err := JobWrapper.CreateJob(ctx, azdClient, jobParams)
=======
			// Show spinner while creating job
			spinner := ux.NewSpinner(&ux.SpinnerOptions{
				Text: "creating fine-tuning job...",
			})
			if err := spinner.Start(ctx); err != nil {
				fmt.Printf("failed to start spinner: %v\n", err)
			}

			// Parse and validate the YAML configuration file if provided
			var config *models.CreateFineTuningRequest
			if filename != "" {
				color.Green("\nparsing configuration file...")
				config, err = utils.ParseCreateFineTuningRequestConfig(filename)
				if err != nil {
					_ = spinner.Stop(ctx)
					fmt.Println()
					return err
				}
			} else {
				config = &models.CreateFineTuningRequest{}
			}

			// Override config values with command-line parameters if provided
			if model != "" {
				config.BaseModel = model
			}
			if trainingFile != "" {

				config.TrainingFile = trainingFile
			}
			if validationFile != "" {
				config.ValidationFile = &validationFile
			}
			if suffix != "" {
				config.Suffix = &suffix
			}
			if seed != 0 {
				config.Seed = &seed
			}

			fineTuneSvc, err := services.NewFineTuningService(ctx, azdClient, nil)
			if err != nil {
				_ = spinner.Stop(ctx)
				fmt.Println()
				return err
			}

			// Submit the fine-tuning job using CreateJob from JobWrapper
			job, err := fineTuneSvc.CreateFineTuningJob(ctx, config)
			_ = spinner.Stop(ctx)
			fmt.Println()

>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
			if err != nil {
				return err
			}

			// Print success message
<<<<<<< HEAD
			fmt.Println(strings.Repeat("=", 120))
			color.Green("\nSuccessfully submitted fine-tuning Job!\n")
			fmt.Printf("Job ID:     %s\n", job.Id)
			fmt.Printf("Model:      %s\n", job.Model)
=======
			fmt.Println("\n", strings.Repeat("=", 60))
			color.Green("\nsuccessfully submitted fine-tuning Job!\n")
			fmt.Printf("Job ID:     %s\n", job.ID)
			fmt.Printf("Model:      %s\n", job.BaseModel)
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
			fmt.Printf("Status:     %s\n", job.Status)
			fmt.Printf("Created:    %s\n", job.CreatedAt)
			if job.FineTunedModel != "" {
				fmt.Printf("Fine-tuned: %s\n", job.FineTunedModel)
			}
<<<<<<< HEAD
			fmt.Println(strings.Repeat("=", 120))

=======
			fmt.Println(strings.Repeat("=", 60))
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
			return nil
		},
	}

<<<<<<< HEAD
	cmd.Flags().StringVarP(&filename, "file", "f", "", "Path to the config file")

	return cmd
}

func newOperationShowCommand() *cobra.Command {
	var jobID string

	cmd := &cobra.Command{
		Use:   "show",
		Short: "Show the fine tuning job details",
=======
	cmd.Flags().StringVarP(&filename, "file", "f", "", "Path to the config file.")
	cmd.Flags().StringVarP(&model, "model", "m", "", "Base model to fine-tune. Overrides config file. Required if --file is not provided")
	cmd.Flags().StringVarP(&trainingFile, "training-file", "t", "", "Training file ID or local path. Use 'local:' prefix for local paths. Required if --file is not provided")
	cmd.Flags().StringVarP(&validationFile, "validation-file", "v", "", "Validation file ID or local path. Use 'local:' prefix for local paths.")
	cmd.Flags().StringVarP(&suffix, "suffix", "s", "", "An optional string of up to 64 characters that will be added to your fine-tuned model name. Overrides config file.")
	cmd.Flags().Int64VarP(&seed, "seed", "r", 0, "Random seed for reproducibility of the job. If a seed is not specified, one will be generated for you. Overrides config file.")

	//Either config file should be provided or at least `model` & `training-file` parameters
	cmd.MarkFlagFilename("file", "yaml", "yml")
	cmd.MarkFlagsOneRequired("file", "model")
	cmd.MarkFlagsRequiredTogether("model", "training-file")
	return cmd
}

// newOperationShowCommand creates a command to show the fine-tuning job details
func newOperationShowCommand() *cobra.Command {
	var jobID string
	var logs bool
	var output string

	cmd := &cobra.Command{
		Use:   "show",
		Short: "Shows detailed information about a specific job.",
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := azdext.WithAccessToken(cmd.Context())
			azdClient, err := azdext.NewAzdClient()
			if err != nil {
				return fmt.Errorf("failed to create azd client: %w", err)
			}
			defer azdClient.Close()
<<<<<<< HEAD
			// Show spinner while fetching jobs
			spinner := ux.NewSpinner(&ux.SpinnerOptions{
				Text: fmt.Sprintf("Fetching fine-tuning job %s...", jobID),
			})
			if err := spinner.Start(ctx); err != nil {
				fmt.Printf("Failed to start spinner: %v\n", err)
			}

			// Fetch fine-tuning job details using job wrapper
			job, err := JobWrapper.GetJobDetails(ctx, azdClient, jobID)
			_ = spinner.Stop(ctx)

			if err != nil {
				return fmt.Errorf("failed to get fine-tuning job details: %w", err)
			}

			// Print job details
			color.Green("\nFine-Tuning Job Details\n")
			fmt.Printf("Job ID:              %s\n", job.Id)
			fmt.Printf("Status:              %s %s\n", getStatusSymbol(job.Status), job.Status)
			fmt.Printf("Model:               %s\n", job.Model)
			fmt.Printf("Fine-tuned Model:    %s\n", formatFineTunedModel(job.FineTunedModel))
			fmt.Printf("Created At:          %s\n", job.CreatedAt)
			if job.FinishedAt != "" {
				fmt.Printf("Finished At:         %s\n", job.FinishedAt)
			}
			fmt.Printf("Method:              %s\n", job.Method)
			fmt.Printf("Training File:       %s\n", job.TrainingFile)
			if job.ValidationFile != "" {
				fmt.Printf("Validation File:     %s\n", job.ValidationFile)
			}

			// Print hyperparameters if available
			if job.Hyperparameters != nil {
				fmt.Println("\nHyperparameters:")
				fmt.Printf("  Batch Size:              %d\n", job.Hyperparameters.BatchSize)
				fmt.Printf("  Learning Rate Multiplier: %f\n", job.Hyperparameters.LearningRateMultiplier)
				fmt.Printf("  N Epochs:                %d\n", job.Hyperparameters.NEpochs)
			}

			// Fetch and print events
			eventsSpinner := ux.NewSpinner(&ux.SpinnerOptions{
				Text: "Fetching job events...",
			})
			if err := eventsSpinner.Start(ctx); err != nil {
				fmt.Printf("Failed to start spinner: %v\n", err)
			}

			events, err := JobWrapper.GetJobEvents(ctx, azdClient, jobID)
			_ = eventsSpinner.Stop(ctx)

			if err != nil {
				fmt.Printf("Warning: failed to fetch job events: %v\n", err)
			} else if events != nil && len(events.Data) > 0 {
				fmt.Println("\nJob Events:")
				for i, event := range events.Data {
					fmt.Printf("  %d. [%s] %s - %s\n", i+1, event.Level, event.CreatedAt, event.Message)
				}
				if events.HasMore {
					fmt.Println("  ... (more events available)")
				}
			}

			// Fetch and print checkpoints if job is completed
			if job.Status == "succeeded" {
				checkpointsSpinner := ux.NewSpinner(&ux.SpinnerOptions{
					Text: "Fetching job checkpoints...",
				})
				if err := checkpointsSpinner.Start(ctx); err != nil {
					fmt.Printf("Failed to start spinner: %v\n", err)
				}

				checkpoints, err := JobWrapper.GetJobCheckPoints(ctx, azdClient, jobID)
				_ = checkpointsSpinner.Stop(ctx)

				if err != nil {
					fmt.Printf("Warning: failed to fetch job checkpoints: %v\n", err)
				} else if checkpoints != nil && len(checkpoints.Data) > 0 {
					fmt.Println("\nJob Checkpoints:")
					for i, checkpoint := range checkpoints.Data {
						fmt.Printf("  %d. Checkpoint ID: %s\n", i+1, checkpoint.ID)
						fmt.Printf("     Checkpoint Name:       %s\n", checkpoint.FineTunedModelCheckpoint)
						fmt.Printf("     Created On:            %s\n", checkpoint.CreatedAt)
						fmt.Printf("     Step Number:           %d\n", checkpoint.StepNumber)
						if checkpoint.Metrics != nil {
							fmt.Printf("     Full Validation Loss:  %.6f\n", checkpoint.Metrics.FullValidLoss)
						}
					}
					if checkpoints.HasMore {
						fmt.Println("  ... (more checkpoints available)")
=======

			// Show spinner while fetching job
			spinner := ux.NewSpinner(&ux.SpinnerOptions{
				Text: "Fine-Tuning Job Details",
			})
			if err := spinner.Start(ctx); err != nil {
				fmt.Printf("failed to start spinner: %v\n", err)
			}

			fineTuneSvc, err := services.NewFineTuningService(ctx, azdClient, nil)
			if err != nil {
				_ = spinner.Stop(ctx)
				fmt.Print("\n\n")
				return err
			}

			job, err := fineTuneSvc.GetFineTuningJobDetails(ctx, jobID)
			_ = spinner.Stop(ctx)
			fmt.Print("\n\n")
			if err != nil {
				return err
			}

			switch output {
			case "json":
				utils.PrintObject(job, utils.FormatJSON)
			case "yaml":
				utils.PrintObject(job, utils.FormatYAML)
			case "table", "":
				views := job.ToDetailViews()
				indent := "  "
				utils.PrintObjectWithIndent(views.Details, utils.FormatTable, indent)

				fmt.Println("\nTimestamps:")
				utils.PrintObjectWithIndent(views.Timestamps, utils.FormatTable, indent)
				fmt.Println("\nConfiguration:")
				utils.PrintObjectWithIndent(views.Configuration, utils.FormatTable, indent)

				fmt.Println("\nData:")
				utils.PrintObjectWithIndent(views.Data, utils.FormatTable, indent)
			default:
				return fmt.Errorf("unsupported output format: %s (supported: table, json, yaml)", output)
			}

			if logs {
				fmt.Println()
				// Fetch and print events
				eventsSpinner := ux.NewSpinner(&ux.SpinnerOptions{
					Text: "Events:",
				})
				if err := eventsSpinner.Start(ctx); err != nil {
					fmt.Printf("failed to start spinner: %v\n", err)
				}

				events, err := fineTuneSvc.GetJobEvents(ctx, jobID)
				_ = eventsSpinner.Stop(ctx)
				fmt.Println()

				if err != nil {
					return err
				} else if events != nil && len(events.Data) > 0 {
					const eventIndent = "     "
					for _, event := range events.Data {
						fmt.Printf("%s[%s] %s\n", eventIndent, utils.FormatTime(event.CreatedAt), event.Message)
					}
					if events.HasMore {
						fmt.Println("  ... (more events available)")
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
					}
				}
			}

<<<<<<< HEAD
			fmt.Println(strings.Repeat("=", 120))

			return nil
		},
	}
	cmd.Flags().StringVarP(&jobID, "job-id", "i", "", "Fine-tuning job ID")
	cmd.MarkFlagRequired("job-id")
	return cmd
}

func newOperationListCommand() *cobra.Command {
	var top int
	var after string
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List the fine tuning jobs",
=======
			return nil
		},
	}

	cmd.Flags().StringVarP(&jobID, "id", "i", "", "Job ID")
	cmd.Flags().BoolVar(&logs, "logs", false, "Include recent training logs")
	cmd.Flags().StringVarP(&output, "output", "o", "table", "Output format: table, json, yaml")
	cmd.MarkFlagRequired("id")

	return cmd
}

// newOperationListCommand creates a command to list fine-tuning jobs
func newOperationListCommand() *cobra.Command {
	var limit int
	var after string
	var output string
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List fine-tuning jobs.",
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := azdext.WithAccessToken(cmd.Context())
			azdClient, err := azdext.NewAzdClient()
			if err != nil {
				return fmt.Errorf("failed to create azd client: %w", err)
			}
			defer azdClient.Close()

			// Show spinner while fetching jobs
			spinner := ux.NewSpinner(&ux.SpinnerOptions{
<<<<<<< HEAD
				Text: "Fetching fine-tuning jobs...",
			})
			if err := spinner.Start(ctx); err != nil {
				fmt.Printf("Failed to start spinner: %v\n", err)
			}

			// List fine-tuning jobs using job wrapper
			jobs, err := JobWrapper.ListJobs(ctx, azdClient, top, after)
			_ = spinner.Stop(ctx)

			if err != nil {
				return fmt.Errorf("failed to list fine-tuning jobs: %w", err)
			}

			for i, job := range jobs {
				fmt.Printf("\n%d. Job ID: %s | Status: %s %s | Model: %s | Fine-tuned: %s | Created: %s",
					i+1, job.Id, getStatusSymbol(job.Status), job.Status, job.Model, formatFineTunedModel(job.FineTunedModel), job.CreatedAt)
			}

			fmt.Printf("\nTotal jobs: %d\n", len(jobs))

			return nil
		},
	}
	cmd.Flags().IntVarP(&top, "top", "t", 50, "Number of fine-tuning jobs to list")
	cmd.Flags().StringVarP(&after, "after", "a", "", "Cursor for pagination")
	return cmd
}

func newOperationActionCommand() *cobra.Command {
	var jobID string
	var action string

	cmd := &cobra.Command{
		Use:   "action",
		Short: "Perform an action on a fine-tuning job (pause, resume, cancel)",
=======
				Text: "Fine-Tuning Jobs",
			})
			if err := spinner.Start(ctx); err != nil {
				fmt.Printf("failed to start spinner: %v\n", err)
			}

			fineTuneSvc, err := services.NewFineTuningService(ctx, azdClient, nil)
			if err != nil {
				_ = spinner.Stop(ctx)
				fmt.Println()
				return err
			}

			jobs, err := fineTuneSvc.ListFineTuningJobs(ctx, limit, after)
			_ = spinner.Stop(ctx)
			fmt.Print("\n\n")

			if err != nil {
				return err
			}

			switch output {
			case "json":
				utils.PrintObject(jobs, utils.FormatJSON)
			case "table", "":
				utils.PrintObject(jobs, utils.FormatTable)
			default:
				return fmt.Errorf("unsupported output format: %s (supported: table, json)", output)
			}
			return nil
		},
	}

	cmd.Flags().IntVarP(&limit, "top", "t", 10, "Number of jobs to return")
	cmd.Flags().StringVar(&after, "after", "", "Pagination cursor")
	cmd.Flags().StringVarP(&output, "output", "o", "table", "Output format: table, json")
	return cmd
}

// newOperationPauseCommand creates a command to pause a running fine-tuning job
func newOperationPauseCommand() *cobra.Command {
	var jobID string

	cmd := &cobra.Command{
		Use:   "pause",
		Short: "Pauses a running fine-tuning job.",
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := azdext.WithAccessToken(cmd.Context())
			azdClient, err := azdext.NewAzdClient()
			if err != nil {
				return fmt.Errorf("failed to create azd client: %w", err)
			}
			defer azdClient.Close()

<<<<<<< HEAD
			// Validate job ID is provided
			if jobID == "" {
				return fmt.Errorf("job-id is required")
			}

			// Validate action is provided and valid
			if action == "" {
				return fmt.Errorf("action is required (pause, resume, or cancel)")
			}

			action = strings.ToLower(action)
			if action != "pause" && action != "resume" && action != "cancel" {
				return fmt.Errorf("invalid action '%s'. Allowed values: pause, resume, cancel", action)
			}

			var job *JobWrapper.JobContract
			var err2 error

			// Execute the requested action
			switch action {
			case "pause":
				job, err2 = JobWrapper.PauseJob(ctx, azdClient, jobID)
			case "resume":
				job, err2 = JobWrapper.ResumeJob(ctx, azdClient, jobID)
			case "cancel":
				job, err2 = JobWrapper.CancelJob(ctx, azdClient, jobID)
			}

			if err2 != nil {
				return err2
			}

			// Print success message
			fmt.Println()
			fmt.Println(strings.Repeat("=", 120))
			color.Green(fmt.Sprintf("\nSuccessfully %sd fine-tuning Job!\n", action))
			fmt.Printf("Job ID:     %s\n", job.Id)
			fmt.Printf("Model:      %s\n", job.Model)
			fmt.Printf("Status:     %s %s\n", getStatusSymbol(job.Status), job.Status)
			fmt.Printf("Created:    %s\n", job.CreatedAt)
			if job.FineTunedModel != "" {
				fmt.Printf("Fine-tuned: %s\n", job.FineTunedModel)
			}
			fmt.Println(strings.Repeat("=", 120))

			return nil
		},
	}

	cmd.Flags().StringVarP(&jobID, "job-id", "i", "", "Fine-tuning job ID")
	cmd.Flags().StringVarP(&action, "action", "a", "", "Action to perform: pause, resume, or cancel")
	cmd.MarkFlagRequired("job-id")
	cmd.MarkFlagRequired("action")

	return cmd
}

func newOperationDeployModelCommand() *cobra.Command {
	var jobID string
	var deploymentName string
	var modelFormat string
	var sku string
	var version string
	var capacity int32

	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy a fine-tuned model to Azure Cognitive Services",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := azdext.WithAccessToken(cmd.Context())
			azdClient, err := azdext.NewAzdClient()
			if err != nil {
				return fmt.Errorf("failed to create azd client: %w", err)
			}
			defer azdClient.Close()

			// Validate required parameters
			if jobID == "" {
				return fmt.Errorf("job-id is required")
			}
			if deploymentName == "" {
				return fmt.Errorf("deployment-name is required")
			}

			// Get environment values
			envValueMap := make(map[string]string)
			if envResponse, err := azdClient.Environment().GetCurrent(ctx, &azdext.EmptyRequest{}); err == nil {
				env := envResponse.Environment
				envValues, err := azdClient.Environment().GetValues(ctx, &azdext.GetEnvironmentRequest{
					Name: env.Name,
				})
				if err != nil {
					return fmt.Errorf("failed to get environment values: %w", err)
				}

				for _, value := range envValues.KeyValues {
					envValueMap[value.Key] = value.Value
				}
			}

			// Create deployment configuration
			deployConfig := JobWrapper.DeploymentConfig{
				JobID:             jobID,
				DeploymentName:    deploymentName,
				ModelFormat:       modelFormat,
				SKU:               sku,
				Version:           version,
				Capacity:          capacity,
				SubscriptionID:    envValueMap["AZURE_SUBSCRIPTION_ID"],
				ResourceGroup:     envValueMap["AZURE_RESOURCE_GROUP_NAME"],
				AccountName:       envValueMap["AZURE_ACCOUNT_NAME"],
				TenantID:          envValueMap["AZURE_TENANT_ID"],
				WaitForCompletion: true,
			}

			// Deploy the model using the wrapper
			result, err := JobWrapper.DeployModel(ctx, azdClient, deployConfig)
=======
			// Show spinner while pausing job
			spinner := ux.NewSpinner(&ux.SpinnerOptions{
				Text: "Pausing fine-tuning job...",
			})
			if err := spinner.Start(ctx); err != nil {
				fmt.Printf("failed to start spinner: %v\n", err)
			}

			fineTuneSvc, err := services.NewFineTuningService(ctx, azdClient, nil)
			if err != nil {
				_ = spinner.Stop(ctx)
				fmt.Println()
				return err
			}

			job, err := fineTuneSvc.PauseJob(ctx, jobID)
			_ = spinner.Stop(ctx)
			fmt.Println()

>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
			if err != nil {
				return err
			}

			// Print success message
<<<<<<< HEAD
			fmt.Println(strings.Repeat("=", 120))
			color.Green("\nSuccessfully deployed fine-tuned model!\n")
			fmt.Printf("Deployment Name:  %s\n", result.DeploymentName)
			fmt.Printf("Status:           %s\n", result.Status)
			fmt.Printf("Message:          %s\n", result.Message)
			fmt.Println(strings.Repeat("=", 120))

=======
			fmt.Println("✓ Job pause request submitted successfully")
			fmt.Println()
			fmt.Printf("  Job ID:  %s\n", job.ID)
			fmt.Printf("  Status:  %s\n", job.Status)
			fmt.Println()
			fmt.Printf("Resume with: azd ai finetune jobs resume --id %s\n", job.ID)
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7
			return nil
		},
	}

<<<<<<< HEAD
	cmd.Flags().StringVarP(&jobID, "job-id", "i", "", "Fine-tuning job ID")
	cmd.Flags().StringVarP(&deploymentName, "deployment-name", "d", "", "Deployment name")
	cmd.Flags().StringVarP(&modelFormat, "model-format", "m", "OpenAI", "Model format")
	cmd.Flags().StringVarP(&sku, "sku", "s", "Standard", "SKU for deployment")
	cmd.Flags().StringVarP(&version, "version", "v", "1", "Model version")
	cmd.Flags().Int32VarP(&capacity, "capacity", "c", 1, "Capacity for deployment")
	cmd.MarkFlagRequired("job-id")
	cmd.MarkFlagRequired("deployment-name")
=======
	cmd.Flags().StringVarP(&jobID, "id", "i", "", "Job ID")
	cmd.MarkFlagRequired("id")

	return cmd
}

// newOperationResumeCommand creates a command to resume a paused fine-tuning job
func newOperationResumeCommand() *cobra.Command {
	var jobID string

	cmd := &cobra.Command{
		Use:   "resume",
		Short: "Resumes a paused fine-tuning job.",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := azdext.WithAccessToken(cmd.Context())
			azdClient, err := azdext.NewAzdClient()
			if err != nil {
				return fmt.Errorf("failed to create azd client: %w", err)
			}
			defer azdClient.Close()

			// Show spinner while resuming job
			spinner := ux.NewSpinner(&ux.SpinnerOptions{
				Text: "Resuming fine-tuning job...",
			})
			if err := spinner.Start(ctx); err != nil {
				fmt.Printf("failed to start spinner: %v\n", err)
			}

			fineTuneSvc, err := services.NewFineTuningService(ctx, azdClient, nil)
			if err != nil {
				_ = spinner.Stop(ctx)
				fmt.Println()
				return err
			}

			job, err := fineTuneSvc.ResumeJob(ctx, jobID)
			_ = spinner.Stop(ctx)
			fmt.Println()

			if err != nil {
				return err
			}

			// Print success message
			fmt.Println("✓ Job resume request submitted successfully")
			fmt.Println()
			fmt.Printf("  Job ID:  %s\n", job.ID)
			fmt.Printf("  Status:  %s\n", job.Status)
			fmt.Println()
			fmt.Printf("View progress: azd ai finetune jobs show --id %s\n", job.ID)
			return nil
		},
	}

	cmd.Flags().StringVarP(&jobID, "id", "i", "", "Job ID")
	cmd.MarkFlagRequired("id")

	return cmd
}

// newOperationCancelCommand creates a command to cancel a fine-tuning job
func newOperationCancelCommand() *cobra.Command {
	var jobID string
	var force bool

	cmd := &cobra.Command{
		Use:   "cancel",
		Short: "Cancels a running or queued fine-tuning job.",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := azdext.WithAccessToken(cmd.Context())
			azdClient, err := azdext.NewAzdClient()
			if err != nil {
				return fmt.Errorf("failed to create azd client: %w", err)
			}
			defer azdClient.Close()

			// Prompt for confirmation unless --force is specified
			if !force {
				fmt.Printf("Cancel fine-tuning job %s? (y/N): ", jobID)
				var response string
				fmt.Scanln(&response)
				response = strings.ToLower(strings.TrimSpace(response))
				if response != "y" && response != "yes" {
					fmt.Println("Operation aborted.")
					return nil
				}
			}

			// Show spinner while canceling job
			spinner := ux.NewSpinner(&ux.SpinnerOptions{
				Text: "Cancelling fine-tuning job...",
			})
			if err := spinner.Start(ctx); err != nil {
				fmt.Printf("failed to start spinner: %v\n", err)
			}

			fineTuneSvc, err := services.NewFineTuningService(ctx, azdClient, nil)
			if err != nil {
				_ = spinner.Stop(ctx)
				fmt.Println()
				return err
			}

			job, err := fineTuneSvc.CancelJob(ctx, jobID)
			_ = spinner.Stop(ctx)
			fmt.Println()

			if err != nil {
				return err
			}

			// Print success message
			fmt.Println("✓ Job cancel request submitted successfully")
			fmt.Println()
			fmt.Printf("  Job ID:  %s\n", job.ID)
			fmt.Printf("  Status:  %s\n", job.Status)
			return nil
		},
	}

	cmd.Flags().StringVarP(&jobID, "id", "i", "", "Job ID")
	cmd.Flags().BoolVar(&force, "force", false, "Skip confirmation prompt")
	cmd.MarkFlagRequired("id")
>>>>>>> 428498d0f124a73e2e722a86cd49d2bf99d05ba7

	return cmd
}
