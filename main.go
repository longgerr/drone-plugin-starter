package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli/v2"
)

var version string // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "my plugin"
	app.Usage = "my plugin usage"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{

		//
		// repo args
		//

		&cli.StringFlag{
			Name:    "repo.fullname",
			Usage:   "repository full name",
			EnvVars: []string{"DRONE_REPO"},
		},
		&cli.StringFlag{
			Name:    "repo.owner",
			Usage:   "repository owner",
			EnvVars: []string{"DRONE_REPO_OWNER"},
		},
		&cli.StringFlag{
			Name:    "repo.name",
			Usage:   "repository name",
			EnvVars: []string{"DRONE_REPO_NAME"},
		},
		&cli.StringFlag{
			Name:    "repo.link",
			Usage:   "repository link",
			EnvVars: []string{"DRONE_REPO_LINK"},
		},
		&cli.StringFlag{
			Name:    "repo.avatar",
			Usage:   "repository avatar",
			EnvVars: []string{"DRONE_REPO_AVATAR"},
		},
		&cli.StringFlag{
			Name:    "repo.branch",
			Usage:   "repository default branch",
			EnvVars: []string{"DRONE_REPO_BRANCH"},
		},
		&cli.BoolFlag{
			Name:    "repo.private",
			Usage:   "repository is private",
			EnvVars: []string{"DRONE_REPO_PRIVATE"},
		},
		&cli.BoolFlag{
			Name:    "repo.trusted",
			Usage:   "repository is trusted",
			EnvVars: []string{"DRONE_REPO_TRUSTED"},
		},

		//
		// commit args
		//

		&cli.StringFlag{
			Name:    "remote.url",
			Usage:   "git remote url",
			EnvVars: []string{"DRONE_REMOTE_URL"},
		},
		&cli.StringFlag{
			Name:    "commit.sha",
			Usage:   "git commit sha",
			EnvVars: []string{"DRONE_COMMIT_SHA"},
		},
		&cli.StringFlag{
			Name:    "commit.ref",
			Value:   "refs/heads/master",
			Usage:   "git commit ref",
			EnvVars: []string{"DRONE_COMMIT_REF"},
		},
		&cli.StringFlag{
			Name:    "commit.branch",
			Value:   "master",
			Usage:   "git commit branch",
			EnvVars: []string{"DRONE_COMMIT_BRANCH"},
		},
		&cli.StringFlag{
			Name:    "commit.message",
			Usage:   "git commit message",
			EnvVars: []string{"DRONE_COMMIT_MESSAGE"},
		},
		&cli.StringFlag{
			Name:    "commit.link",
			Usage:   "git commit link",
			EnvVars: []string{"DRONE_COMMIT_LINK"},
		},
		&cli.StringFlag{
			Name:    "commit.author.name",
			Usage:   "git author name",
			EnvVars: []string{"DRONE_COMMIT_AUTHOR"},
		},
		&cli.StringFlag{
			Name:    "commit.author.email",
			Usage:   "git author email",
			EnvVars: []string{"DRONE_COMMIT_AUTHOR_EMAIL"},
		},
		&cli.StringFlag{
			Name:    "commit.author.avatar",
			Usage:   "git author avatar",
			EnvVars: []string{"DRONE_COMMIT_AUTHOR_AVATAR"},
		},

		//
		// build args
		//

		&cli.StringFlag{
			Name:    "build.event",
			Value:   "push",
			Usage:   "build event",
			EnvVars: []string{"DRONE_BUILD_EVENT"},
		},
		&cli.IntFlag{
			Name:    "build.number",
			Usage:   "build number",
			EnvVars: []string{"DRONE_BUILD_NUMBER"},
		},
		&cli.IntFlag{
			Name:    "build.created",
			Usage:   "build created",
			EnvVars: []string{"DRONE_BUILD_CREATED"},
		},
		&cli.IntFlag{
			Name:    "build.started",
			Usage:   "build started",
			EnvVars: []string{"DRONE_BUILD_STARTED"},
		},
		&cli.IntFlag{
			Name:    "build.finished",
			Usage:   "build finished",
			EnvVars: []string{"DRONE_BUILD_FINISHED"},
		},
		&cli.StringFlag{
			Name:    "build.status",
			Usage:   "build status",
			Value:   "success",
			EnvVars: []string{"DRONE_BUILD_STATUS"},
		},
		&cli.StringFlag{
			Name:    "build.link",
			Usage:   "build link",
			EnvVars: []string{"DRONE_BUILD_LINK"},
		},
		&cli.StringFlag{
			Name:    "build.deploy",
			Usage:   "build deployment target",
			EnvVars: []string{"DRONE_DEPLOY_TO"},
		},
		&cli.BoolFlag{
			Name:    "yaml.verified",
			Usage:   "build yaml is verified",
			EnvVars: []string{"DRONE_YAML_VERIFIED"},
		},
		&cli.BoolFlag{
			Name:    "yaml.signed",
			Usage:   "build yaml is signed",
			EnvVars: []string{"DRONE_YAML_SIGNED"},
		},

		//
		// prev build args
		//

		&cli.IntFlag{
			Name:    "prev.build.number",
			Usage:   "previous build number",
			EnvVars: []string{"DRONE_PREV_BUILD_NUMBER"},
		},
		&cli.StringFlag{
			Name:    "prev.build.status",
			Usage:   "previous build status",
			EnvVars: []string{"DRONE_PREV_BUILD_STATUS"},
		},
		&cli.StringFlag{
			Name:    "prev.commit.sha",
			Usage:   "previous build sha",
			EnvVars: []string{"DRONE_PREV_COMMIT_SHA"},
		},
	}

	app.Run(os.Args)
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Repo: Repo{
			Owner:   c.String("repo.owner"),
			Name:    c.String("repo.name"),
			Link:    c.String("repo.link"),
			Avatar:  c.String("repo.avatar"),
			Branch:  c.String("repo.branch"),
			Private: c.Bool("repo.private"),
			Trusted: c.Bool("repo.trusted"),
		},
		Build: Build{
			Number:   c.Int("build.number"),
			Event:    c.String("build.event"),
			Status:   c.String("build.status"),
			Deploy:   c.String("build.deploy"),
			Created:  int64(c.Int("build.created")),
			Started:  int64(c.Int("build.started")),
			Finished: int64(c.Int("build.finished")),
			Link:     c.String("build.link"),
		},
		Commit: Commit{
			Remote:  c.String("remote.url"),
			Sha:     c.String("commit.sha"),
			Ref:     c.String("commit.sha"),
			Link:    c.String("commit.link"),
			Branch:  c.String("commit.branch"),
			Message: c.String("commit.message"),
			Author: Author{
				Name:   c.String("commit.author.name"),
				Email:  c.String("commit.author.email"),
				Avatar: c.String("commit.author.avatar"),
			},
		},
		Config: Config{
			// plugin-specific parameters
		},
	}

	if err := plugin.Exec(); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
