#!/usr/bin/env python3
from github import Github
from datetime import datetime, timedelta

username = 'QuesoTech'
repo_name = 'PhotoHunter'

commit_temp = '- {name}: {message}'
comment_temp = '\t- {name}: {message}'

collaborators = {
        'connorgreenwell': 'connor',
        'emallson': 'david',
        'rbalten': 'ryan',
        'atbradshaw': 'aaron'
        }
if __name__ == '__main__':
    with open('token.txt') as tf:
        token = tf.read().strip()
        github = Github(token, 'x-oauth-basic')

    repo = github.get_user(username).get_repo(repo_name)

    now = datetime.now()
    week = timedelta(days=7)

    print('# Recent Commits')

    for commit in repo.get_commits(since=now-week):
        print(commit_temp.format(message=commit.commit.message,
            name=commit.commit.committer.name))
        comments = commit.get_comments()
        for comment in comments:
            print(comment_temp.format(name=comment.user.name,
                message=comment.body))
