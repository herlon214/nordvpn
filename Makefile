release:
	git tag `cat .version`
	git push --tags