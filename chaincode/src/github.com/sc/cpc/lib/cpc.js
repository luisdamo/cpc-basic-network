/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { Contract } = require('fabric-contract-api');

class Cpc extends Contract {

    async initLedger(ctx) {
        console.info('============= START : Initialize Ledger ===========');
        const assets = [
            {
                DMC: 'DMC00000000001',
                TYPE: 'CART0001',
                ST: 0,
                IDMAN: 'MAN1',
                IDASS: 'ASS1',
                IDCUS: 'CUS1',
            },
            {
                DMC: 'DMC00000000002',
                TYPE: 'GEAR0002',
                ST: 0,
                IDMAN: 'MAN1',
                IDASS: 'ASS1',
                IDCUS: 'CUS1',
            },
            {
                DMC: 'DMC00000000003',
                TYPE: 'OILP0010',
                ST: 0,
                IDMAN: 'MAN1',
                IDASS: 'ASS1',
                IDCUS: 'CUS1',
            },
            {
                DMC: 'DMC00000000004',
                TYPE: 'CART0001',
                ST: 0,
                IDMAN: 'MAN1',
                IDASS: 'ASS1',
                IDCUS: 'CUS1',
            },
            {
                DMC: 'DMC00000000005',
                TYPE: 'CART0002',
                ST: 0,
                IDMAN: 'MAN1',
                IDASS: 'ASS1',
                IDCUS: 'CUS1',
            },
            {
                DMC: 'DMC00000000006',
                TYPE: 'CART0001',
                ST: 0,
                IDMAN: 'MAN1',
                IDASS: 'ASS1',
                IDCUS: 'CUS1',
            },
            {
                DMC: 'DMC00000000007',
                TYPE: 'CART0001',
                ST: 0,
                IDMAN: 'MAN1',
                IDASS: 'ASS1',
                IDCUS: 'CUS1',
            },
            {
                DMC: 'DMC00000000008',
                TYPE: 'CART0001',
                ST: 0,
                IDMAN: 'MAN1',
                IDASS: 'ASS1',
                IDCUS: 'CUS1',
            },
            {
                DMC: 'DMC00000000009',
                TYPE: 'CART0001',
                ST: 0,
                IDMAN: 'MAN1',
                IDASS: 'ASS1',
                IDCUS: 'CUS1',
            },
            {
                DMC: 'DMC00000000010',
                TYPE: 'CART0001',
                ST: 0,
                IDMAN: 'MAN1',
                IDASS: 'ASS1',
                IDCUS: 'CUS1',
            },
        ];

        for (let i = 0; i < assets.length; i++) {
            assets[i].docType = 'car part';
            await ctx.stub.putState('ASSET' + i, Buffer.from(JSON.stringify(assets[i])));
            console.info('Added <--> ', assets[i]);
        }
        console.info('============= END : Initialize Ledger ===========');
    }

    async createAsset(ctx, assetNumber, DMC, TYPE, ST, IDMAN, IDASS, IDCUS) {
        console.info('============= START : Create Asset ===========');

        const asset = {
            docType: 'car part',
            DMC,
            TYPE,
            ST,
            IDMAN,
            IDASS,
            IDCUS,
        };

        await ctx.stub.putState(assetNumber, Buffer.from(JSON.stringify(asset)));
        console.info('============= END : Create Asset ===========');
    }

    async queryAsset(ctx, assetNumber) {
        const assetAsBytes = await ctx.stub.getState(assetNumber); // get the asset from chaincode state
        if (!assetAsBytes || assetAsBytes.length === 0) {
            throw new Error(`${assetNumber} does not exist`);
        }
        console.log(assetAsBytes.toString());
        return assetAsBytes.toString();
    }

    async queryAllAssets(ctx) {
        const startKey = 'ASSET0';
        const endKey = 'ASSET999';

        const iterator = await ctx.stub.getStateByRange(startKey, endKey);

        const allResults = [];
        while (true) {
            const res = await iterator.next();

            if (res.value && res.value.value.toString()) {
                console.log(res.value.value.toString('utf8'));

                const Key = res.value.key;
                let Record;
                try {
                    Record = JSON.parse(res.value.value.toString('utf8'));
                } catch (err) {
                    console.log(err);
                    Record = res.value.value.toString('utf8');
                }
                allResults.push({ Key, Record });
            }
            if (res.done) {
                console.log('end of data');
                await iterator.close();
                console.info(allResults);
                return JSON.stringify(allResults);
            }
        }
    }

    async changeAssetST(ctx, assetNumber, newST) {
        console.info('============= START : changeAssetST ===========');

        const assetAsBytes = await ctx.stub.getState(assetNumber); // get the asset from chaincode state
        if (!assetAsBytes || assetAsBytes.length === 0) {
            throw new Error(`${assetNumber} does not exist`);
        }
        const asset = JSON.parse(assetAsBytes.toString());
        asset.ST = newST;

        await ctx.stub.putState(assetNumber, Buffer.from(JSON.stringify(asset)));
        console.info('============= END : changeAssetST ===========');
    }

}

module.exports = Cpc;
