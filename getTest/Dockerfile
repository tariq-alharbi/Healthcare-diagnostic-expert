FROM public.ecr.aws/lambda/python:3.8

COPY . ./


COPY app.py ${LAMBDA_TASK_ROOT}
COPY model.py ${LAMBDA_TASK_ROOT}
COPY DecisionTree_model.sav ${LAMBDA_TASK_ROOT}
# Install the function's dependencies using file requirements.txt
# from your project folder.

COPY requirements.txt .
RUN  pip3 install -r requirements.txt --target "${LAMBDA_TASK_ROOT}"

# Command can be overwritten by providing a different command in the template directly.
CMD ["app.lambda_handler"]
