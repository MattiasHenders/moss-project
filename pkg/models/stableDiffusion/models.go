package stableDiffusion

type Request struct {
	URL          string
	Method       string
	Body         []byte
	OkStatusCode int
	Headers      map[string]string
	Data         string
	Bearer       string
}

type Response struct {
	Body       []byte
	StatusCode int
}

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
	Prompt           string   `json:"prompt"`
	Seed             *int64   `json:"seed,omitempty"`
	NumOutputs       *int     `json:"batch_size,omitempty"`
	Width            *int     `json:"width,omitempty"`
	Height           *int     `json:"height,omitempty"`
	NumInfrenceSteps *int     `json:"steps,omitempty"`
	GuidanceScale    *float64 `json:"guidance_scale,omitempty"`
	InitImage        *int     `json:"init_image,omitempty"`
	Strength         *float64 `json:"strength,omitempty"`
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
