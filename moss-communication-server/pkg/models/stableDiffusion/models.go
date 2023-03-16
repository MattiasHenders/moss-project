package stableDiffusion

type StableDiffusionRequest struct {
	Input StableDiffusionInput `json:"input"`
}

type StableDiffusionResponse struct {
	ID            string                `json:"id"`
	Input         StableDiffusionInput  `json:"input"`
	Status        string                `json:"status"`
	DelayTime     int64                 `json:"delayTime,omitempty"`
	ExecutionTime int64                 `json:"executionTime,omitempty"`
	Output        StableDiffusionOutput `json:"output,omitempty"`
}

type StableDiffusionInput struct {
	Prompt string `json:"prompt"`
}

type StableDiffusionOutput struct {
	Images     []string                  `json:"images,omitempty"`
	Info       string                    `json:"info,omitempty"`
	Parameters StableDiffusionParameters `json:"parameters,omitempty"`
}

type StableDiffusionParameters struct {
	BatchSize         int    `json:"batch_size,omitempty"`
	CFGScle           int    `json:"cfg_scale,omitempty"`
	DenoisingStrength int    `json:"denoising_strength,omitempty"`
	Height            int    `json:"height,omitempty"`
	Width             int    `json:"width,omitempty"`
	NIter             int    `json:"n_iter,omitempty"`
	Prompt            string `json:"prompt,omitempty"`
	RestoreFaces      bool   `json:"restore_faces,omitempty"`
	SChurn            int    `json:"s_churn,omitempty"`
	SNoise            int    `json:"s_noise,omitempty"`
	STMin             int    `json:"s_tmin,omitempty"`
	SamplerIndex      string `json:"sampler_index,omitempty"`
	Seed              int    `json:"seed,omitempty"`
	Steps             int    `json:"steps,omitempty"`
	Tiling            bool   `json:"tiling,omitempty"`
}
