from joblib import load
model = load('DecisionTree_model.sav')

def predict (data):
    data=list(data)
    yhat = model.predict([data])
    data=(str(yhat[0]).replace("[",'')).replace("]",'')
    data=data.replace('.','').split()
    for c in data:
        if len(c)>1:
            l=list(c)
            data[data.index(c)]=str(l[0])+"."+str(l[1])
       
    return data







