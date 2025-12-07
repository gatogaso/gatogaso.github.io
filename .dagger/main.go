// A generated module for Bulsara functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/bulsara/internal/dagger"
)

type Bulsara struct{}

// Returns a basic hugo environment container.
func (m *Bulsara) BuildBasicHugoEnv(ctx context.Context, hugoVersion string, directoryArg *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("hugomods/hugo:exts-"+hugoVersion).
		WithMountedDirectory("/src", directoryArg).
		WithWorkdir("/src")
}

// Save your progress online
func (m *Bulsara) CommitAndPush(ctx context.Context, directoryArg *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("commitizen/commitizen:latest").
		WithMountedDirectory("/src", directoryArg).
		WithWorkdir("/src").
		WithExec([]string{"git", "config", "core.quotepath", "false"}).
		WithExec([]string{"git", "add", "content/*", "static/*", "hugo.toml"}).
		WithExec([]string{"echo", "'Write: cz c here below and proceed.'"})// .
		// WithExec([]string{"cz", "c"}).
		// WithExec([]string{"git", "push"})
}

// TODO: Remove if not needed
// // Bump version of the project and print it
// func (m *Bulsara) TagProject(ctx context.Context, directoryArg *dagger.Directory) *dagger.Container {
// 	return dag.Container().
// 		From("commitizen/commitizen:latest").
// 		WithMountedDirectory("/src", directoryArg).
// 		WithWorkdir("/src").
// 		WithExec([]string{"cz", "bump"})
// }

// Return a container with hugo server running. Adding: | up --ports=1313:1313
// E.g.: dagger -c 'develop-hugo 0.152.2 . | up --ports=1313:1313'
func (m *Bulsara) DevelopHugo(ctx context.Context, hugoVersion string, directoryArg *dagger.Directory) *dagger.Service {
	return m.BuildBasicHugoEnv(ctx, hugoVersion, directoryArg).
		WithExposedPort(1313).
		AsService(dagger.ContainerAsServiceOpts{Args: []string{"hugo", "server", "--bind", "0.0.0.0", "--port", "1313", "--buildDrafts", "--buildFuture", "--disableFastRender", "--renderToMemory", "--poll", "700ms", "--noHTTPCache", "--gc", "-w"}})
}

// Returns the web statics compiled with hugo. Needs to add | export ./public
// E.g.: dagger -c 'build-hugo 0.152.2 . test.com /tmp | export ./public'
func (m *Bulsara) BuildHugo(ctx context.Context, hugoVersion string, directoryArg *dagger.Directory, baseURL string, tempCachePath string) *dagger.Directory {
	// outputs := dag.Directory()

	builder := m.BuildBasicHugoEnv(ctx, hugoVersion, directoryArg).
		WithExec([]string{"hugo", "--gc", "--minify", "--baseURL", baseURL, "--cacheDir", tempCachePath+"/hugo_cache"})

	// return outputs.WithDirectory("./public", builder.Directory("./public"))
	return builder.Directory("./public")
}

// Executes terraform apply to deploy the infrastructure
func (m *Bulsara) DeployInfra(ctx context.Context, directoryArg *dagger.Directory) (string, error) {
	// TODO(dmoreno): Add the WithSecretVariable for the AWS keys so the can be invoked like --secret-access-key env://AWS_SECRET_ACCESS_KEY
	// TODO(dmoreno): Adapt the terraform command properly
	return dag.Container().
		From("hashicorp/terraform").
		WithMountedDirectory("/src", directoryArg).
		WithWorkdir("/src").
		WithExec([]string{"sh", "-c", "terraform -chdir=infrastructure init"}).
		WithExec([]string{"sh", "-c", "terraform -chdir=infrastructure apply -auto-approve -no-color"}).
		Stdout(ctx)
}

// Executes terraform destroy to remove the infrastructure
func (m *Bulsara) DestroyInfra(ctx context.Context, directoryArg *dagger.Directory) (string, error) {
	// TODO(dmoreno): Add the WithSecretVariable for the AWS keys so the can be invoked like --secret-access-key env://AWS_SECRET_ACCESS_KEY
	// TODO(dmoreno): Adapt the terraform command properly
	return dag.Container().
		From("hashicorp/terraform").
		WithMountedDirectory("/src", directoryArg).
		WithWorkdir("/src").
		WithExec([]string{"sh", "-c", "terraform -chdir=infrastructure init"}).
		WithExec([]string{"sh", "-c", "terraform -chdir=infrastructure destroy -auto-approve -no-color"}).
		Stdout(ctx)
}
