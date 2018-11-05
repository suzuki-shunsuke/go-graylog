if [ -n "$CIRCLE_PULL_REQUEST" ]; then
  npx commitlint --from master --to $CIRCLE_BRANCH || exit 1
elif [ "$CIRCLE_BRANCH" != "master" ]; then
  npx commitlint --from master --to HEAD || exit 1
else
  npx commitlint --from HEAD~1 --to HEAD || exit 1
fi
