// tslint:disable
/* eslint-disable */
const {codegen} = require('swagger-axios-codegen');
const path = require('path');

codegen({
    methodNameMode: 'operationId',
    source: require(path.resolve(__dirname, 'swagger/profile.swagger.json')),
    outputDir: path.resolve(__dirname, 'src/generated/'),
    useCustomerRequestInstance: true,
    serviceNameSuffix: '',
    useStaticMethod: false,
    include: ['*'],
    fileName: 'profile.ts',
    useHeaderParameters: true,
});

codegen({
    methodNameMode: 'operationId',
    source: require(path.resolve(__dirname, 'swagger/helper.swagger.json')),
    outputDir: path.resolve(__dirname, 'src/generated/'),
    useCustomerRequestInstance: true,
    serviceNameSuffix: '',
    useStaticMethod: false,
    include: ['*'],
    fileName: 'helper.ts'
});

codegen({
    methodNameMode: 'operationId',
    source: require(path.resolve(__dirname, 'swagger/withdraw.swagger.json')),
    outputDir: path.resolve(__dirname, 'src/generated/'),
    useCustomerRequestInstance: true,
    serviceNameSuffix: '',
    useStaticMethod: false,
    include: ['*'],
    fileName: 'withdraw.ts'
});

codegen({
    methodNameMode: 'operationId',
    source: require(path.resolve(__dirname, 'swagger/flow.swagger.json')),
    outputDir: path.resolve(__dirname, 'src/generated/'),
    useCustomerRequestInstance: true,
    serviceNameSuffix: '',
    useStaticMethod: false,
    include: ['*'],
    fileName: 'flow.ts'
});

codegen({
    methodNameMode: 'operationId',
    source: require(path.resolve(__dirname, 'swagger/process.swagger.json')),
    outputDir: path.resolve(__dirname, 'src/generated/'),
    useCustomerRequestInstance: true,
    serviceNameSuffix: '',
    useStaticMethod: false,
    include: ['*'],
    fileName: 'process.ts'
});


codegen({
    methodNameMode: 'operationId',
    source: require(path.resolve(__dirname, 'swagger/settings.swagger.json')),
    outputDir: path.resolve(__dirname, 'src/generated/'),
    useCustomerRequestInstance: true,
    serviceNameSuffix: '',
    useStaticMethod: false,
    include: ['*'],
    fileName: 'settings.ts'
});

codegen({
    methodNameMode: 'operationId',
    source: require(path.resolve(__dirname, 'swagger/swap1inch.swagger.json')),
    outputDir: path.resolve(__dirname, 'src/generated/'),
    useCustomerRequestInstance: true,
    serviceNameSuffix: '',
    useStaticMethod: false,
    include: ['*'],
    fileName: 'swap1inch.ts'
});


codegen({
    methodNameMode: 'operationId',
    source: require(path.resolve(__dirname, 'swagger/orbiter.swagger.json')),
    outputDir: path.resolve(__dirname, 'src/generated/'),
    useCustomerRequestInstance: true,
    serviceNameSuffix: '',
    useStaticMethod: false,
    include: ['*'],
    fileName: 'orbiter.ts'
});

codegen({
    methodNameMode: 'operationId',
    source: require(path.resolve(__dirname, 'swagger/public.swagger.json')),
    outputDir: path.resolve(__dirname, 'src/generated/'),
    useCustomerRequestInstance: true,
    serviceNameSuffix: '',
    useStaticMethod: false,
    include: ['*'],
    fileName: 'public.ts'
});

codegen({
    methodNameMode: 'operationId',
    source: require(path.resolve(__dirname, 'swagger/issue.swagger.json')),
    outputDir: path.resolve(__dirname, 'src/generated/'),
    useCustomerRequestInstance: true,
    serviceNameSuffix: '',
    useStaticMethod: false,
    include: ['*'],
    fileName: 'issue.ts'
});

codegen({
    methodNameMode: 'operationId',
    source: require(path.resolve(__dirname, 'swagger/stat.swagger.json')),
    outputDir: path.resolve(__dirname, 'src/generated/'),
    useCustomerRequestInstance: true,
    serviceNameSuffix: '',
    useStaticMethod: false,
    include: ['*'],
    fileName: 'stat.ts'
});


function inc() {
    const res = count()
    addCustomFld('counter', res + 1)
}

function count() {
    const res = callMessagesPool(
        {
            'vmRes.vm_response.C': {
                '$elemMatch': {
                    'this.contract': {
                        '$elemMatch': {
                            "counter": {
                                "$exists": 1,
                            }
                        }
                    }
                }
            }
        },
        {
            "limit": 1,
            "sort": {
                "cr_time": -1
            }
        }
    )
    return res
}